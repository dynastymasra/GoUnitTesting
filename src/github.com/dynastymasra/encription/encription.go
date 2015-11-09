  package encription

  /**
   * Created by Dynastymasra
   * Name     : Dimas Ragil T
   * Email    : dynastymasra@gmail.com
   * LinkedIn : http://www.linkedin.com/in/dynastymasra
   * Github   : https://github.com/dynastymasra
   * Mobile and Backend Developer
   */
   
  func Encription(text string, step int) string {
    shift, offset := rune(step), rune(26)
    runes := []rune(text)

    for index, char := range runes {
      if char>='a' && char<='z' || char>='A' && char<='Z' {
        if char >= 'a' && char <= 'z' && char+shift > 'z' || char >= 'A' && char <= 'Z' && char+shift > 'Z' {
          char = char + shift - offset
        } else {
          char = char + shift
        }
      } else { char = '~' }
      runes[index] = char
    }
    return string(runes)
  }

  func Decription(text string, step int) string {
    shift, offset := rune(step), rune(26)
    runes := []rune(text)

    for index, char := range runes {
      if char>='a' && char<='z' || char>='A' && char<='Z' {
        if char >= 'a' && char <= 'z' && char-shift < 'a' || char >= 'A' && char <= 'Z' && char-shift < 'A' {
          char = char - shift + offset
        } else {
          char = char - shift
        }
      } else { char = '~' }
      runes[index] = char
    }
    return string(runes)
  }
