package main

import "github.com/ArkjuniorK/go-crypto-cli/cmd"

func main() {
	cmd.Run()

	//key := "THISISMYKEYWORDPASS"
	//msg := "THIS IS WHY USING GO LANGUAGE IS MORE CREATIVE"
	//
	//v := cipher.NewVigenere()
	//{
	//	enc := v.Encipher(msg, key)
	//	dec := v.Decipher(enc, key)
	//
	//	fmt.Println("Vigenere Cipher")
	//	fmt.Println("Encoded", enc)
	//	fmt.Println("Decoded", dec)
	//}
	//
	//fmt.Println("--------------------------------")
	//
	//p := cipher.NewPlayfair(key)
	//{
	//	var m = []byte(msg)
	//
	//	encb := make([]byte, len(m))
	//	p.Encrypt(encb, m)
	//	enc := bytes.Buffer{}
	//	enc.Write(encb)
	//
	//	decb := make([]byte, len([]byte(enc.String())))
	//	p.Decrypt(decb, []byte(enc.String()))
	//	dec := bytes.Buffer{}
	//	dec.Write(decb)
	//
	//	fmt.Println("Playfair Cipher")
	//	fmt.Println("Encoded", enc.String())
	//	fmt.Println("Decoded", dec.String())
	//}
	//
	//fmt.Println("--------------------------------")
	//
	//m, _ := enigmamachine.New(enigmamachine.MachineSetup{
	//	Reflector:     reflectors.B,
	//	Rotors:        []enigmamachine.RotorSpec{rotors.I, rotors.II, rotors.III},
	//	RingPositions: []int{10, 14, 21},
	//	Plugboard:     "AP BR CM FZ GJ IL NT OV QS WX",
	//})
	//{
	//	fmt.Println("Enigma Cipher")
	//
	//	m.SetPositions([]rune{'V', 'Q', 'Q'})
	//	result, _ := m.TranslateString(msg)
	//	fmt.Println("Encode", result)
	//
	//	m.SetPositions([]rune{'V', 'Q', 'Q'})
	//	result, _ = m.TranslateString(result)
	//	fmt.Println("Decode", result)
	//
	//}
	//
	//fmt.Println("--------------------------------")
	//
	//mt := make([]byte, 16*2048)
	//rand.Read(mt)
	//o, _ := otp.NewPad(mt, 100, 1)
	//{
	//	encrypted, err := o.Encrypt([]byte(msg))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Encode", base64.StdEncoding.EncodeToString(encrypted))
	//
	//	decrypted, err := o.Decrypt(encrypted)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("Decode %s\n", decrypted)
	//}

}
