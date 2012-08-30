/* Matrix package with emphasis on speed
 *
 *	All operations are in-place on receiver:
 *
 *		a := Zeros()
 *		b := Ones()
 *		a.Add(b)
 *
 *	- All constructors return objects, not interfaces
 *	- All methods accept read-only Matrix interface, except specialized
 *		methods like AddDense, AddSparse
 *	- Explicit dimensions are avoided to allow Zeros, Ones, Eye, etc to have any size 
 */
package mat

type Matrix interface {
	At(i, j int) float64
}

type Rows interface {
	Rows() int
}

type Cols interface {
	Cols() int
}

type RowsMatrix interface {
	Rows
	Matrix
}

type ColsMatrix interface {
	Cols
	Matrix
}

type RowsColsMatrix interface {
	Matrix
	Rows
	Cols
}
