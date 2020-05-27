var digits  map[byte]int
func init () {
    digits = map[byte]int{
		'0':0,
		'1':1,
		'2':2,
		'3':3,
		'4':4,
		'5':5,
		'6':6,
		'7':7,
		'8':8,
		'9':9,
	}
}
var max int = 2<<30 -1
var min int = -2<<30
func myAtoi(s string) (res int) {
    
    i:=0
    sign := 1
    for ; i < len (s); i++ {
        if s[i] == ' '{
            continue
        }
        if s[i] == '-' {
            sign = -1
            i++ 
            break
        } 
        if s[i] == '+' {
            i++
            break
        }
        if s[i] >= '0' || s[i] <= '9' {
            break
        }
        return 
    }
    for ; i < len (s); i++ {
        if s[i] < '0' || s[i] > '9' {
            break
        }
        res = res * 10 + digits[s[i]] 
        if res * sign> max {
           return max
         }
         if res * sign < min {
           return min
          }
    }
    res = res * sign
      
    return
}
