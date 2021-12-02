package uuid

import (
	. "gopkg.in/check.v1"
)

type genTestSuite struct{}

var _ = Suite(&genTestSuite{})

func (s *genTestSuite) TestNewV1(c *C) {
	u := NewV1()
	c.Assert(u.Version(), Equals, V1)
	c.Assert(u.Variant(), Equals, VariantRFC4122)

	u1 := NewV1()
	u2 := NewV1()
	c.Assert(u1, Not(Equals), u2)

	oldFunc := epochFunc
	epochFunc = func() uint64 { return 0 }

	u3 := NewV1()
	u4 := NewV1()
	c.Assert(u3, Not(Equals), u4)

	epochFunc = oldFunc
}

func (s *genTestSuite) BenchmarkNewV1(c *C) {
	for i := 0; i < c.N; i++ {
		NewV1()
	}
}

func (s *genTestSuite) TestNewV2(c *C) {
	u1 := NewV2(DomainPerson)
	c.Assert(u1.Version(), Equals, V2)
	c.Assert(u1.Variant(), Equals, VariantRFC4122)

	u2 := NewV2(DomainGroup)
	c.Assert(u2.Version(), Equals, V2)
	c.Assert(u2.Variant(), Equals, VariantRFC4122)
}

func (s *genTestSuite) BenchmarkNewV2(c *C) {
	for i := 0; i < c.N; i++ {
		NewV2(DomainPerson)
	}
}

func (s *genTestSuite) TestNewV3(c *C) {
	u := NewV3(NamespaceDNS, "www.example.com")
	c.Assert(u.Version(), Equals, V3)
	c.Assert(u.Variant(), Equals, VariantRFC4122)
	c.Assert(u.String(), Equals, "5df41881-3aed-3515-88a7-2f4a814cf09e")

	u = NewV3(NamespaceDNS, "python.org")
	c.Assert(u.String(), Equals, "6fa459ea-ee8a-3ca4-894e-db77e160355e")

	u1 := NewV3(NamespaceDNS, "golang.org")
	u2 := NewV3(NamespaceDNS, "golang.org")
	c.Assert(u1, Equals, u2)

	u3 := NewV3(NamespaceDNS, "example.com")
	c.Assert(u1, Not(Equals), u3)

	u4 := NewV3(NamespaceURL, "golang.org")
	c.Assert(u1, Not(Equals), u4)
}

func (s *genTestSuite) BenchmarkNewV3(c *C) {
	for i := 0; i < c.N; i++ {
		NewV3(NamespaceDNS, "www.example.com")
	}
}

func (s *genTestSuite) TestNewV4(c *C) {
	u := NewV4()
	c.Assert(u.Version(), Equals, V4)
	c.Assert(u.Variant(), Equals, VariantRFC4122)
}

func (s *genTestSuite) BenchmarkNewV4(c *C) {
	for i := 0; i < c.N; i++ {
		NewV4()
	}
}

func (s *genTestSuite) TestNewV5(c *C) {
	u := NewV5(NamespaceDNS, "www.example.com")
	c.Assert(u.Version(), Equals, V5)
	c.Assert(u.Variant(), Equals, VariantRFC4122)

	u = NewV5(NamespaceDNS, "python.org")
	c.Assert(u.String(), Equals, "886313e1-3b8a-5372-9b90-0c9aee199e5d")

	u1 := NewV5(NamespaceDNS, "golang.org")
	u2 := NewV5(NamespaceDNS, "golang.org")
	c.Assert(u1, Equals, u2)

	u3 := NewV5(NamespaceDNS, "example.com")
	c.Assert(u1, Not(Equals), u3)

	u4 := NewV5(NamespaceURL, "golang.org")
	c.Assert(u1, Not(Equals), u4)
}

func (s *genTestSuite) BenchmarkNewV5(c *C) {
	for i := 0; i < c.N; i++ {
		NewV5(NamespaceDNS, "www.example.com")
	}
}
