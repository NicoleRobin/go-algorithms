package segment_tree

/*
SegmentTree segment tree
data: origin number array
tree: tree implementation with array
n: number array len
treeType: tree type
*/
type SegmentTree struct {
	data     []int
	tree     []int
	n        int
	treeType SegmentTreeType
}

type SegmentTreeType string

const (
	SegmentTreeTypeMax SegmentTreeType = "max"
	SegmentTreeTypeMin SegmentTreeType = "min"
	SegmentTreeTypeSum SegmentTreeType = "sum"
)

// NewSegmentTree 构建线段树
func NewSegmentTree(treeType SegmentTreeType, nums []int) *SegmentTree {
	n := len(nums)
	tree := make([]int, 4*n) // 线段树大小一般取 4 倍空间
	st := &SegmentTree{
		treeType: treeType,
		data:     nums,
		tree:     tree,
		n:        n,
	}
	st.build(1, 0, n-1)
	return st
}

/*
build 构建线段树，其实就是在tree中计算对应区间的值
node: 树节点，1是根节点，左孩子节点是node*2，右孩子节点是node*2+1
l, r: 是区间，当l==r时表示到了叶子节点
*/
func (st *SegmentTree) build(node, l, r int) {
	if l == r {
		st.tree[node] = st.data[l]
		return
	}
	mid := (l + r) / 2
	st.build(node*2, l, mid)
	st.build(node*2+1, mid+1, r)
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

// Query 区间和查询
func (st *SegmentTree) Query(qL, qR int) int {
	return st.query(1, 0, st.n-1, qL, qR)
}

/*
node: 树节点
l, r: node对应的区间
qL, qR: 查询区间
*/
func (st *SegmentTree) query(node, l, r, qL, qR int) int {
	if qL > r || qR < l {
		return 0 // 不相交
	}
	if qL <= l && r <= qR {
		return st.tree[node] // 完全包含
	}
	mid := (l + r) / 2
	left := st.query(node*2, l, mid, qL, qR)
	right := st.query(node*2+1, mid+1, r, qL, qR)
	return left + right
}

// Update 更新某个位置的值
func (st *SegmentTree) Update(index, val int) {
	st.update(1, 0, st.n-1, index, val)
}

/*
update 更新指定的值
node: 树节点
l, r: node对应的区间
index: 要更新值的下标
val: 要更新的值
*/
func (st *SegmentTree) update(node, l, r, index, val int) {
	if l == r {
		st.tree[node] = val
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		st.update(node*2, l, mid, index, val)
	} else {
		st.update(node*2+1, mid+1, r, index, val)
	}
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}
