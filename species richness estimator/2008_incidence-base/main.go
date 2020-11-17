package main 

import(
	"gonum.org/v1/gonum/optimize"
	"fmt"
	"math"
	"math/rand"
	"bohao"
	
)

type estimator struct{
	T int
	S int 
	individuals int 
	t int

	
	sample_list [][]int

	D int
	D_list []int

	fk_list []int 


	p_list  []float64
}
func C_(a, b int) int{
	return bohao.Factorial(int(a)) / bohao.Factorial(int(a-b)) / bohao.Factorial(int(b))
}

func k(alpha, beta float64, T int) (K float64){
	term1 := math.Gamma(alpha) * (math.Gamma(beta) / math.Gamma(alpha + beta))
    term2 := math.Gamma(alpha) * (math.Gamma(beta + float64(T)) / math.Gamma(alpha + beta + float64(T))) 

	K = 1 / (term1 - term2)
	return K
}

func (E *estimator) cal_p(x []float64) (p_list []float64){
	alpha, beta := x[0], x[1]
	K := k(x[0], x[1], E.T)
	for x := 0.; x <= float64(E.t); x ++{
		var p float64
		if x == 0.{
			term1 := math.Gamma(alpha) * math.Gamma(float64(E.t) + beta) / math.Gamma(float64(E.t) + alpha + beta)
			term2 := math.Gamma(alpha) * math.Gamma(float64(E.T) + beta) / math.Gamma(float64(E.T) + alpha + beta)
			p = K * (term1 - term2)
		}else {
			p = K * float64(C_(E.t, int(x))) *math.Gamma(alpha + x)*math.Gamma(float64(E.t) + beta - x)/ math.Gamma(float64(E.t) + alpha + beta)
		}
		p_list = append(p_list, p)
	}	
	
	return p_list 
}

func (E *estimator) Ln_L(x []float64) float64{
	res := 0.
	E.p_list = E.cal_p(x)

	for k := 0; k < E.t; k++{
		res += float64(E.fk_list[k]) * (math.Log(E.p_list[k + 1]) - math.Log(1. - E.p_list[0]))
	}
	return -res 
}


func (E *estimator) Init() {
	for i := 0; i < E.t; i++{
		var tmp []int 
		for j := 0; j < E.individuals; j ++{
			tmp = append(tmp, rand.Intn(E.S))
		}
		E.sample_list = append(E.sample_list, tmp)
	}

	
	for i := 0; i < len(E.sample_list); i ++{
		E.D_list = append(E.D_list, E.sample_list[i]...)
	}
	E.D_list = bohao.RemoveDuplicateElement_int(E.D_list)
	E.D = len(E.D_list)

	fk_list := make([]int, E.t)
	for i := 0; i < E.D; i ++{
		f := 0
		for j := 0; j < E.t; j ++{
			if bohao.IsIn_int(E.D_list[i], E.sample_list[j]){
				f += 1
			}
		}
		fk_list[f - 1] += 1
		
	}
	E.fk_list = fk_list
}

func (E *estimator) Find_S(){
	

	p := optimize.Problem{
		Func: E.Ln_L,
	}
	
	x0 := []float64{1, 1}
	result, err := optimize.Minimize(p, x0, nil, nil)
	if err != nil {
		panic(err)
	}
	if err = result.Status.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("result.Status: %v\n", result.Status)
	fmt.Printf("result.X: %0.4g\n", result.X)
	fmt.Printf("result.F: %0.4g\n", result.F)
	fmt.Printf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations)

}




func main() {

	E := estimator{
		T: 100, 
		S: 110,
		t: 10, 
		individuals: 20, 
	}

	E.Init()
	E.Find_S()

	
	
	/* p := optimize.Problem{
		Func: f,
	}
	
	x := []float64{1.3, 1.2}
	result, err := optimize.Minimize(p, x, nil, nil)
	if err != nil {
		panic(err)
	}
	if err = result.Status.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("result.Status: %v\n", result.Status)
	fmt.Printf("result.X: %0.4g\n", result.X)
	fmt.Printf("result.F: %0.4g\n", result.F)
	fmt.Printf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations) */
	
}