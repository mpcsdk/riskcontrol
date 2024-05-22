package pendding

import (
	"context"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/logic/tfa/verifier"
	"riskcontral/internal/service"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *RiskPendding) TfaTidyMail(ctx context.Context, userId string, mail string) error {
	riskKind := s.riskKind
	if riskKind != tfaconst.RiskKind_BindMail && riskKind != tfaconst.RiskKind_UpMail {
		return nil
	}
	if mail == "" {
		return mpccode.CodeParamInvalid()
	}
	notexists, err := service.DB().TfaDB().TfaMailNotExists(ctx, mail)
	if err != nil || !notexists {
		return mpccode.CodeTFAMailExists()
	}
	////
	if riskKind == tfaconst.RiskKind_BindMail {
		return s.tfaBindMail(ctx, userId, mail)
	} else {
		return s.tfaUpMail(ctx, userId, mail)
	}
}

// ////
// ///
func (s *RiskPendding) tfaBindMail(ctx context.Context, userId string, mail string) error {

	verifier := verifier.NewVerifierMail(tfaconst.RiskKind_BindMail, mail)

	s.AddVerifier(verifier)
	s.AddAfterFunc(func(ctx context.Context) error {
		err := recordMail(ctx, userId, mail, false)
		return err
	})

	return nil
}

func (s *RiskPendding) tfaUpMail(ctx context.Context, userId string, mail string) error {
	verifier := verifier.NewVerifierMail(tfaconst.RiskKind_UpMail, mail)

	s.AddVerifier(verifier)
	s.AddAfterFunc(func(ctx context.Context) error {
		err := recordMail(ctx, userId, mail, true)
		return err
	})

	//
	return nil
	//
}
