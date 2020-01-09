package been

import (
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/bn256"
)

func generate() (*bn256.G2, error) {
	secret, err := rand.Int(rand.Reader, bn256.Order)
	if err != nil {
		return nil, err
	}
	return new(bn256.G2).ScalarBaseMult(secret), nil
}

func writeAndLoad(g2 *bn256.G2) (*bn256.G2, error) {
	result := new(bn256.G2)
	_, err := result.Unmarshal(g2.Marshal())
	return result, err
}

func Xenophobic() error {
	g2, err := generate()
	if err != nil {
		return err
	}
	g22, err := writeAndLoad(g2)
	if err != nil {
		return err
	}
	fmt.Println(g22.Marshal())
	fmt.Println(g2.Marshal())
	return nil
}
