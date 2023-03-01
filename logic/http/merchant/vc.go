package merchant

import (
	"net/http"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/did"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

func GetMerchantVCList(r *http_util.HTTPContext) (resp interface{}, err error) {
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	return db.FindPaginationMerchantVC(jwtClaims.MerchantId, db.VCStatusCreated, (page-1)*size, size)
}

type MarkVCReceivedRequest struct {
	VcIds []string `json:"vc_ids"`
}

func MarkVCReceived(req *MarkVCReceivedRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	return nil, db.MarkVCReceived(jwtClaims.MerchantId, req.VcIds)
}

type MarkAllVCInvalidRequest struct {
	VcIds []string `json:"vc_ids"`
	All   bool     `json:"all"`
}

func MarkAllVCInvalid(req *MarkAllVCInvalidRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	return nil, db.WithTx(func(tx db.Options) error {
		for {
			var (
				vcs []*db.VC
				err error
			)

			if req.All {
				vcs, err = db.FindPaginationMerchantVC(jwtClaims.MerchantId, db.VCStatusActive, 0, 100, tx, db.WithPreload("Payment.Currency.Blockchain"))
				if err != nil {
					return err
				}
			} else {
				vcs, err = db.FindPaginationMerchantVCByVCId(jwtClaims.MerchantId, req.VcIds, db.VCStatusActive, 0, 100, tx, db.WithPreload("Payment.Currency.Blockchain"))
				if err != nil {
					return err
				}
			}

			if len(vcs) == 0 {
				return nil
			}

			for _, vc := range vcs {
				vc.VCID = util.RandString(32)

				vcDocument := did.CreateVerifiableCredential(vc.VCID, did.IssueDID, &did.VCSubjectDeposit{
					Chain:             vc.Payment.Currency.Blockchain.ContractName,
					Currency:          vc.Payment.Currency.Symbol,
					Amount:            vc.Payment.CollectionAmount.String(),
					PlatformFeeAmount: vc.Payment.PlatformFeeAmount.String(),
					PoolFeeAmount:     vc.Payment.PoolFeeAmount.String(),
					MerchantAmount:    vc.Payment.CollectionAmount.Copy().Sub(vc.Payment.PoolFeeAmount).Sub(vc.Payment.PlatformFeeAmount).String(),
				})

				err = vcDocument.Signature(did.IssuePriKey)
				if err != nil {
					return err
				}

				vc.VCContent = vcDocument.ToJson()
				vc.VCStatus = db.VCStatusCreated

				err = db.UpdateVC(vc, db.VCStatusActive, tx)
				if err != nil {
					return err
				}
			}
		}
	})
}

func BatchGetVcStatus(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*batchGetPaymentVcStatusRequest)

	if len(req.VcIds) == 0 || len(req.VcIds) > 500 {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid vc id count")
	}

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// init vc res
	vcSet := make(map[string]*paymentVCItem, 0)
	for _, vcId := range req.VcIds {
		vcSet[vcId] = &paymentVCItem{
			VCID:     vcId,
			VCStatus: "NotFound",
		}
	}

	// get vc status
	vcs, vcErr := db.FindVCsByVCId(req.VcIds, db.WithSelect([]string{"vc_id", "vc_status"}))
	if vcErr != nil {
		return nil, vcErr
	}

	for _, vcItem := range vcs {
		vcSet[vcItem.VCID].VCStatus = vcItem.VCStatus
	}

	paymentVCSet := make([]paymentVCItem, 0)
	for _, item := range vcSet {
		paymentVCSet = append(paymentVCSet, *item)
	}

	return &batchGetPaymentVcStatusResponse{
		Total: uint(len(vcSet)),
		Vcs:   paymentVCSet,
	}, nil
}
