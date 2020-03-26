/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: calculator.i

package calculator

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_calculator_b52f4b2bb402add4(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_calculator_b52f4b2bb402add4(swig_intgo arg1);
extern uintptr_t _wrap_new_Calculator_calculator_b52f4b2bb402add4(swig_intgo arg1, swig_intgo arg2);
extern void _wrap_delete_Calculator_calculator_b52f4b2bb402add4(uintptr_t arg1);
extern swig_intgo _wrap_Calculator_Sum_calculator_b52f4b2bb402add4(uintptr_t arg1);
extern swig_intgo _wrap_Calculator_Sub_calculator_b52f4b2bb402add4(uintptr_t arg1);
extern swig_intgo _wrap_Calculator_Factorial_calculator_b52f4b2bb402add4(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_calculator_b52f4b2bb402add4(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_calculator_b52f4b2bb402add4(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrCalculator uintptr

func (p SwigcptrCalculator) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCalculator) SwigIsCalculator() {
}

func NewCalculator(arg1 int, arg2 int) (_swig_ret Calculator) {
	var swig_r Calculator
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Calculator)(SwigcptrCalculator(C._wrap_new_Calculator_calculator_b52f4b2bb402add4(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func DeleteCalculator(arg1 Calculator) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Calculator_calculator_b52f4b2bb402add4(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrCalculator) Sum() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Calculator_Sum_calculator_b52f4b2bb402add4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCalculator) Sub() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Calculator_Sub_calculator_b52f4b2bb402add4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCalculator) Factorial() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Calculator_Factorial_calculator_b52f4b2bb402add4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

type Calculator interface {
	Swigcptr() uintptr
	SwigIsCalculator()
	Sum() (_swig_ret int)
	Sub() (_swig_ret int)
	Factorial() (_swig_ret int)
}


