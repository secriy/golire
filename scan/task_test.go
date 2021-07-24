package scan

// func TestScan_GetAllPorts(t *testing.T) {
// 	type fields struct {
// 		IPPool   []string
// 		PortPool []uint16
// 		Process  int
// 	}
// 	type args struct {
// 		port string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{"1", fields{[]string{}, []uint16{}, 0}, args{"1,2,2-12"}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &Scan{
// 				IPPool:   tt.fields.IPPool,
// 				PortPool: tt.fields.PortPool,
// 				Process:  tt.fields.Process,
// 			}
// 			if err := s.GetAllPorts(tt.args.port); (err != nil) != tt.wantErr {
// 				t.Errorf("GetAllPorts() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			log.Println(s.PortPool)
// 		})
// 	}
// }
