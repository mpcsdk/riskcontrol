package tfa

import (
	"context"
	"riskcontral/internal/model"
	"testing"
)

var userId = "userId"
var riskSerial = "riskSerial"
var phone = "13812345678"
var s *sTFA = &sTFA{
	riskPenddingContainer: newRiskPenddingContainer(10),
}

func Test_riskPhone(t *testing.T) {

	risk := s.riskPenddingContainer.NewRiskPendding(userId, riskSerial, RiskKind_Tx)
	verifier := newVerifierPhone(RiskKind_Tx, phone)
	risk.AddBeforFunc(nil)
	risk.AddAfterFunc(nil)
	///
	verifier.SetCode("123")
	////verify, failed
	k, err := risk.VerifierCode(&model.VerifyCode{
		PhoneCode: "123",
	})
	if err != nil {
		t.Log(k)
		t.Error(err)
	}
	if k, err = risk.AllDone(); err != nil {
		t.Log(k)
		t.Error(err)
	}

	// ///
	k, err = risk.DoFunc(context.TODO())
	if err != nil {
		t.Error(k)
		t.Error(err)
	}
}

// func Test_riskPhoneMail(t *testing.T) {
// 	var s *sTFA = &sTFA{
// 		riskPenddingContainer: newRiskPenddingContainer(10),
// 	}

// 	///add a phoneriskevent
// 	event := newRiskEventPhone(userId, func(ctx context.Context) error {
// 		t.Log("newRiskEventPhone")
// 		return nil
// 	})
// 	s.riskPenddingContainer.Add(userId, riskSerial, event)
// 	///add mail
// 	event = newRiskEventMail(userId, func(ctx context.Context) error {
// 		t.Log("newRiskEventMail")
// 		return nil
// 	})
// 	s.riskPenddingContainer.Add(userId, riskSerial, event)
// 	////
// 	///, up verificationcode
// 	s.riskPenddingContainer.UpCode(userId, riskSerial, Key_RiskEventPhone, "123")
// 	s.riskPenddingContainer.UpCode(userId, riskSerial, Key_RiskEventMail, "456")
// 	////verify, failed
// 	k, err := s.riskPenddingContainer.VerifierCode(userId, riskSerial, &model.VerifyCode{
// 		PhoneCode: "456",
// 	})
// 	if err != nil {
// 		t.Log(k)
// 		t.Error(err)
// 	}
// 	if k, err = s.riskPenddingContainer.AllDone(userId, riskSerial); err == nil {
// 		t.Error(s.riskPenddingContainer.AllDone(userId, riskSerial))
// 	}
// 	///
// 	//verify succesc
// 	k, err = s.riskPenddingContainer.VerifierCode(userId, riskSerial, &model.VerifyCode{
// 		PhoneCode: "123",
// 		MailCode:  "456",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if k, err := s.riskPenddingContainer.AllDone(userId, riskSerial); err != nil {
// 		t.Log(k)
// 		t.Error(s.riskPenddingContainer.AllDone(userId, riskSerial))
// 	}
// 	///
// 	err = s.riskPenddingContainer.DoAfter(context.TODO(), userId, riskSerial)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func Test_riskTimeOut(t *testing.T) {
// 	var s *sTFA = &sTFA{
// 		riskPenddingContainer: newRiskPenddingContainer(10),
// 	}

// 	///add a phoneriskevent
// 	event := newRiskEventPhone(userId, func(ctx context.Context) error {
// 		t.Log("newRiskEventPhone")
// 		return nil
// 	})
// 	s.riskPenddingContainer.Add(userId, riskSerial, event)
// 	///, up verificationcode
// 	s.riskPenddingContainer.UpCode(userId, riskSerial, Key_RiskEventPhone, "123")
// 	///
// 	r := s.riskPenddingContainer.Get(userId, riskSerial)
// 	if r == nil {
// 		t.Error(r)
// 	}
// 	///
// 	time.Sleep(10 * time.Second)
// 	///
// 	r = s.riskPenddingContainer.Get(userId, riskSerial)
// 	if r != nil {
// 		t.Error(r)
// 	}
// }
