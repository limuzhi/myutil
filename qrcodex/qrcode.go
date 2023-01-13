/*
 * @PackageName: qrcodex
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 17:26
 */

package qrcodex

import (
	"fmt"
	"github.com/boombuler/barcode/qr"
)

type QrOptions struct {
	Width      int                     //生成图片宽度
	Height     int                     //生成图片高度
	Quality    qr.ErrorCorrectionLevel //图片质量 0-100
	LogoSwitch bool                    //是否生成水印logo
	LogoWidth  int                     //水印logo宽度
	LogoHeight int                     //水印logo高度
}

// Option gorm option interface
type OptionFunc interface {
	Apply(*QrOptions) error
	Initialize()
}

// Apply update config to new config
func (o *QrOptions) Apply(cfg *QrOptions) error {
	if cfg != o {
		*cfg = *o
	}
	return nil
}

func NewQrOptions(opts ...OptionFunc) (*QrOptions, error) {
	o := &QrOptions{
		Width:      256,
		Height:     256,
		Quality:    qr.M,
		LogoSwitch: false,
		LogoWidth:  20,
		LogoHeight: 20,
	}
	for _, opt := range opts {
		if opt != nil {
			if applyErr := opt.Apply(o); applyErr != nil {
				return nil, applyErr
			}
			defer func(opt OptionFunc) {
				if errr := opt.AfterInitialize(db); errr != nil {
					err = errr
				}
			}(opt)
		}
	}
	fmt.Println(o)
	return o, nil
}

//func (o *QrOptions) GenerateQrCode(url string) (string, error) {
//	qrCode, err := qr.Encode(url, o.quality, qr.Auto)
//	if err != nil {
//		return "", err
//	}
//
//	qrCode, err = barcode.Scale(qrCode, o.width, o.height)
//	if err != nil {
//		return "", err
//	}
//}
