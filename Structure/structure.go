package structure

type BoxXoso struct {
	Xoso []Xoso `json:"box_kq"`
}
type Xoso struct {
	Day  string `json:"day"`
	Kqxs Kqxs   `json:"kqxs"`
}
type Kqxs struct {
	DB   string   `json:"db"`
	Nhat string   `json:"nhat"`
	Nhi  []string `json:"nhi"`
	Ba   []string `json:"ba"`
	Bon  []string `json:"bon"`
	Nam  []string `json:"nam"`
	Sau  []string `json:"sau"`
	Bay  []string `json:"bay"`
}
