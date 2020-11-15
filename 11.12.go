package main //p295 3DES
import(
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)
func main() {
	key := []byte("abcdefghijklmnopqrstuvwx")//密钥占24个字节
	fmt.Println("---------3DES加密解密字节数组")
	str := "我爱GO语言"
	result, _ := TripleDesEncrypt([]byte(str), key)
	fmt.Printf("加密后: %x\n", result)
	origData, _ := TripleDesDecrypt(result, key)
	fmt.Println("解密后:", string(origData))
	fmt.Println("-------3DES加密解密字符串")
	cipherText, _ := TripleDesEncrypt2Str(str, key)
	fmt.Println("加密后: ",cipherText)
	originalText, _ := TripleDesDecrypt2Str(cipherText, key)
	fmt.Println("解密后:", originalText)
}
//3DES加密字节数组，返回字节数组
func TripleDesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err!= nil {
		return nil, err
	}
	originalBytes = PKCS5Padding(originalBytes, block.BlockSize)
	// originalBytes = ZeroPadding(originalBytes, block.BlockSize())
	blockMide := cipher.NewCBCEncrypter(block, key[:8])
	cipherArr := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(cipherArr, originalBytes)
	return cipherArr, nil
}
//3DES解密字节数组,返回字节数组
func TripleDesDecrypt(cipherBytes, key []byte) ([]byte, error) {
	block, err := des,NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	orignalArr := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(originalArr, cipherBytes)
	originalArr = PKCS5UnPadding(originalArr, cipherBytes)
	// origData = ZeroUnPadding(origData)
	return originalArr, nil
}
//3DES加密字符串，返回base64处理后字符串
func TripleDesEncrypt2Str(originalText string, key []byte) (string, error) {
	block, err := des,NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	originalData := PKCS5Padding([]byte(originalText), block.BlockSize())
	// originalData = ZeroPadding(originalData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherArr := make([]byte, len(originalData))
	blockMode.CryptBlocks(cipherArr, originalData)
	cipherText := base64.StdEncoding.EncodeToString(ciherarr)
	return cipherText, nil
}
// 3DES解密base64处理后的加密字符串，返回明文字符串
func TripleDesDecrypt2Str(cipherText string, key []byte) (string, error) {
	cipherArr, _ := base64.StdEncoding.DecodeSting(cipherText)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	originalArr := make([]byte, len(cipherArr))
	blockMode.CryptBlocks(originalArr, cipherArr)
	originalArr = PKCs5UnPadding(originaArr)
	// origData = ZeroUnPadding(origData)
	return string(originalArr), nil
}
func PkCs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}