package pendding

import (
	"context"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/logic/tfa/verifier"
	"riskcontral/internal/service"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *RiskPendding) TfaTidyPhone(ctx context.Context, userId string, phone string) error {
	////
	if phone == "" {
		return mpccode.CodeParamInvalid()
	}
	notexists, err := service.DB().TfaDB().TfaPhoneNotExists(ctx, phone)
	if err != nil || !notexists {
		return mpccode.CodeTFAPhoneExists()
	}

	///
	if s.riskKind == tfaconst.RiskKind_BindPhone {
		return s.tfaBindPhone(ctx, userId, phone)
	} else {
		return s.tfaUpPhone(ctx, userId, phone)
	}
}

// //
func (s *RiskPendding) tfaBindPhone(ctx context.Context, userId string, phone string) error {
	// if tfaInfo == nil || tfaInfo.Phone != "" {
	// 	return "", mpccode.CodeParamInvalid()
	// }
	verifier := verifier.NewVerifierPhone(tfaconst.RiskKind_BindPhone, phone)
	// risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, tfaconst.RiskKind_BindPhone)

	s.AddVerifier(verifier)
	s.AddAfterFunc(func(ctx context.Context) error {
		err := recordPhone(ctx, userId, phone, false)
		return err
	})

	//
	return nil
}

func (s *RiskPendding) tfaUpPhone(ctx context.Context, userId string, phone string) error {
	// if tfaInfo == nil || tfaInfo.Phone == "" {
	// 	return "", mpccode.CodeParamInvalid()
	// }
	verifier := verifier.NewVerifierPhone(tfaconst.RiskKind_UpPhone, phone)

	s.AddVerifier(verifier)
	s.AddAfterFunc(func(ctx context.Context) error {
		err := recordPhone(ctx, userId, phone, true)
		return err
	})

	//
	return nil
	//
}
