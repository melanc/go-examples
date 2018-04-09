package common
 
import (
    "errors"
    "fmt"
    "github.com/ethereum/go-ethereum/rlp"
)
 
type Args struct {
    A, B float64
    F CBFunc
}
type Result struct {
    Value float64
}
type CBFunc func(s float64)
type MathService struct {
    cb CBFunc
}

func (c Args)MarshalJSON() ([]byte, error) {
    return rlp.EncodeToBytes(c)
}

func (c *Args)UnmarshalJSON(bs []byte) error {
    return rlp.DecodeBytes(bs, *c)
}


func (s *MathService) Add(args *Args, result *Result) error {
    result.Value = args.A + args.B
    fmt.Printf("Add method")
    if s.cb != nil {
        s.cb(result.Value)
    }
    return nil
}
 
func (s *MathService) Divide(args *Args, result *Result) error {
    if args.B == 0 {
        return errors.New("除数不能为零！")
    }
    result.Value = args.A / args.B
    return nil
}

func (s *MathService) SetCB(args *Args, result *Result) error {
    s.cb = args.F
    fmt.Printf("SetCB method\n")
    return nil
}