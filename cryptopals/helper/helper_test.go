package helper

import "testing"

func TestHex(t *testing.T) {
    cases := []struct{in, want string}{
      {"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "I'm killing your brain like a poisonous mushroom"}    }
      for _, c := range cases {
          got := HexToString(c.in)
          if got != c.want{
              t.Errorf("HexToStringError(%q) == %q, want %q", c.in, got, c.want)
          }
      }
}


func Testb64(t *testing.T) {
    cases := []struct{in, want string}{
        {"I'm killing your brain like a poisonous mushroom", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"} }
      for _, c := range cases {
          got := Base64Converter(c.in)
          if got != c.want{
              t.Errorf("Base64Error(%q) == %q, want %q", c.in, got, c.want)
          }
      }
}


func TestXor(t *testing.T) {
    cases := []struct{stringBuf1, stringBuf2, want string}{
        {"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"} }
      for _, c := range cases {
          got := XORStrings(c.stringBuf1, c.stringBuf2)
          if got != c.want{
              t.Errorf("XORError(%q) == %q, want %q", c.in, got, c.want)
          }
      }
}
