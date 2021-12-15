package model

type ContainerRec struct {
	ID        int    `json:"id"`
	HostId    int    `json:"host_id"`
	Name      string `json:"name"`
	ImageName int    `json:"image_name"`
}

type ContainerCreateReq struct {
	HostId    int    `json:"host_id"`
	Name      string `json:"name"`
	ImageName int    `json:"image_name"`
}
