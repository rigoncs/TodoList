package encrypt

import (
	"github.com/rigoncs/TodoList/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type PwdEncryptService struct {
}

const (
	PasswordCost = 12 // 密码加密难度
)

func NewPwdEncryptService() repository.PwdEncrypt {
	return &PwdEncryptService{}
}

func (p *PwdEncryptService) Encrypt(pwd []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pwd, PasswordCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (p *PwdEncryptService) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (p *PwdEncryptService) Check(pwd, src []byte) bool {
	return bcrypt.CompareHashAndPassword(pwd, src) == nil
}
