package hydra

func DefaultPostgresqlList() *AuthList {
	a := NewAuthList()
	a.Username = []string{
		"postgres",
		"test",
		//"user",
		//"root",
		//"manager",
		//"webadmin",
	}
	a.Password = []string{
		"",
		"1",
		"123456",
		"zaq1@WSX",
		"%user%",
		"%user%123",
		"%user%1234",
		"%user%123456",
		"%user%12345",
		"%user%@123",
		"%user%@123456",
		"%user%@12345",
		"%user%#123",
		"%user%#123456",
		"%user%#12345",
		"%user%_123",
		"%user%_123456",
		"%user%_12345",
		"%user%123!@#",
		"%user%!@#$",
		"%user%!@#",
		"%user%~!@",
		"%user%!@#123",
		"%user%2022",
		"%user%2021",
		"%user%2020",
		"%user%2019",
		"%user%2018",
		"%user%2017",
		"%user%2016",
		"%user%2015",
		"%user%@2017",
		"%user%@2016",
		"%user%@2015",
		"qweasdzxc",
		"Passw0rd",
		"admin123!@#",
		"admin",
		"admin123",
		"admin@123",
		"admin#123",
		"password",
		"12345",
		"1234",
		"123",
		"qwerty",
		"1q2w3e4r",
		"1qaz2wsx",
		"qazwsx",
		"123qwe",
		"123qaz",
		"1234567",
		"123456qwerty",
		"password123",
		"12345678",
		"1q2w3e",
		"abc123",
		"test123",
		"123456789",
		"q1w2e3r4",
		//"0000",
		//"root",
		//"test",
		//"okmnji",
		//"postgres",
		//"user",
	}
	a.Special = []Auth{
		//NewSpecialAuth("postgres", "postgres"),
	}
	return a
}
