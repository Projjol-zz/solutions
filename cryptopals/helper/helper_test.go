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
              t.Errorf("HexToStringError(%q) == %q, want %q", c.in, got, c.want)
          }
      }
}
