/* Matrix package with emphasis on speed
 *
 *	All operations are in-place on receiver:
 *
 *		a := mat.New(rows, cols)
 *		a.Add(mat.Ones)
 *		a.Subtract(mat.IDENTITY)
 *		a.Product(mat.ONES, mat.IDENTITY)
 *
 *	- All constructors return objects, not interfaces
 *	- All methods accept read-only Matrix interface, except specialized
 *		methods like AddDense, AddSparse
 *	- Explicit dimensions are avoided to allow Zeros, Ones, Eye, etc to have any size 
 *	- No runtime checks, not even dimensions 
 */
package mat

type Dense struct {
	elems      []float64
	rows, cols int
}

func (m *Dense) At(i, j int) float64 {

	return m.elems[i*m.cols+j]
}

func (m *Dense) Rows() int {
	return m.rows
}

func (m *Dense) Cols() int {
	return m.cols
}

func (m *Dense) Add(a Matrix) {

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.elems[i*m.cols+j] += a.At(i, j)
		}
	}
}

func (m *Dense) Product(a ColsMatrix, b RowsMatrix) {

	n := a.Cols()

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			s := 0.0
			for k := 0; k < n; k++ {
				s += a.At(i, k) * b.At(k, j)
			}
			m.elems[i*m.cols+j] += s
		}
	}
}

func New(rows, cols int) *Dense {

	return &Dense{make([]float64, rows*cols), rows, cols}
}

func (m *Dense) Set(i, j int, v float64) {

	m.elems[i*m.cols+j] = v
}
