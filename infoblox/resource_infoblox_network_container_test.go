package infoblox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var resCfgNetworkContainer_create_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_1" {
  network_view = "%s"
  cidr = "10.0.0.0/16"
  comment = "10.0.0.0/16 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test location"
    Site = "Test site"
  })
}`, testNetView)

var resCfgNetworkContainer_create2_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_2" {
  network_view = "%s"
  cidr = "25.0.0.0/24"
  comment = "25.0.0.0/24 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test location"
    Site = "Test site"
  })
}

resource "infoblox_ipv4_network_container" "nc_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
  allocate_prefix_len = 26
  comment = "dynamic creation of network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}`, testNetView)

var resCfgNetworkContainer_update_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_1" {
  network_view = "%s"
  cidr = "10.0.0.0/16"
  comment = "new 10.0.0.0/16 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test loc. 2"
  })
}`, testNetView)

var resCfgNetworkContainer_update2_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_1" {
  network_view = "%s"
  cidr = "10.0.0.0/16"
  comment = ""
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
  })
}`, testNetView)

var resCfgNetworkContainer_update3_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_2" {
  network_view = "%s"
  cidr = "25.0.0.0/24"
  comment = "25.0.0.0/24 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test location"
    Site = "Test site"
  })
}

resource "infoblox_ipv4_network_container" "nc_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
  allocate_prefix_len = 26
  comment = ""
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
  })
}`, testNetView)

var resCfgNetworkContainer_update4_ipv4 = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc_2" {
  network_view = "%s"
  cidr = "25.0.0.0/24"
  comment = "25.0.0.0/24 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test location"
    Site = "Test site"
  })
}

resource "infoblox_ipv4_network_container" "nc_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
  allocate_prefix_len = 28
  comment = ""
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
  })
}`, testNetView)

var resCfgNetworkContainer_create_ipv6 = fmt.Sprintf(`
resource "infoblox_ipv6_network_container" "nc_1" {
  network_view = "%s"
  cidr = "fc00::/56"
  comment = "fc00::/56 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test loc."
    Site = "Test site"
  })
}`, testNetView)

var resCfgNetworkContainer_create2_ipv6 = fmt.Sprintf(`
resource "infoblox_ipv6_network_container" "nc6_2" {
  network_view = "%s"
  cidr = "fc01::/56"
  comment = "fc01::/56 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "nc6_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
  allocate_prefix_len = 58
  comment = "fc01::/58 dynamic network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}`, testNetView)

var resCfgNetworkContainer_update_ipv6 = fmt.Sprintf(`
resource "infoblox_ipv6_network_container" "nc_1" {
  network_view = "%s"
  cidr = "fc00::/56"
  comment = "new comment for fc00::/56 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Location = "Test loc. 2"
  })
}`, testNetView)

var resCfgNetworkContainer_update2_ipv6 = fmt.Sprintf(`
resource "infoblox_ipv6_network_container" "nc6_2" {
  network_view = "%s"
  cidr = "fc01::/56"
  comment = "fc01::/56 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "nc6_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
  allocate_prefix_len = 58
  comment = "dynamic network container testing"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}`, testNetView)

var resCfgNetworkContainer_update3_ipv6 = fmt.Sprintf(`
resource "infoblox_ipv6_network_container" "nc6_2" {
  network_view = "%s"
  cidr = "fc01::/56"
  comment = "fc01::/56 network container"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "nc6_3" {
  network_view = "default"
  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
  allocate_prefix_len = 59
  comment = "dynamic network container testing"
  ext_attrs = jsonencode({
	"Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}`, testNetView)

func validateNetworkContainer(
	resourceName string,
	expectedValue *ibclient.NetworkContainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id := res.Primary.ID
		if id == "" {
			return fmt.Errorf("ID is not set")
		}
		stateAttr := res.Primary.Attributes["cidr"]

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"terraform_test_tenant")

		nc, err := objMgr.GetNetworkContainerByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}

		expNv := expectedValue.NetviewName
		if nc.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				nc.NetviewName, expNv)
		}

		expComment := expectedValue.Comment
		if nc.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				nc.Comment, expComment)
		}

		expCidr := expectedValue.Cidr
		//cidr is not passed in nextavailable network container test case
		if expCidr == "" {
			expCidr = stateAttr
			if expCidr == "" {
				return fmt.Errorf(
					"the value of 'cidr' field is empty, but expected some value")
			}
		}

		if nc.Cidr != expCidr {
			return fmt.Errorf(
				"the value of 'cidr' field is '%s', but expected '%s'",
				nc.Cidr, expCidr)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && nc.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && nc.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(nc.Ea, expectedEAs)
	}
}

var updateNotAllowedRegexp = regexp.MustCompile("changing the value of '.+' field is not allowed")

func TestAcc_resourceNetworkContainer_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: resCfgNetworkContainer_create_ipv4,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/16",
						Comment:     "10.0.0.0/16 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_update_ipv4,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/16",
						Comment:     "new 10.0.0.0/16 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc. 2",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_update2_ipv4,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/16",
						Comment:     "",
						Ea:          ibclient.EA{},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_create2_ipv4,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_3",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Comment:     "dynamic creation of network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_update3_ipv4,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_3",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Comment:     "",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
						},
					},
				),
			},
			{
				Config:      resCfgNetworkContainer_update4_ipv4,
				ExpectError: updateNotAllowedRegexp,
			},
		},
	})
}

func TestAcc_resourceNetworkContainer_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: resCfgNetworkContainer_create_ipv6,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Cidr:        "fc00::/56",
						Comment:     "fc00::/56 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_update_ipv6,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Cidr:        "fc00::/56",
						Comment:     "new comment for fc00::/56 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc. 2",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_create2_ipv6,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_3",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Comment:     "fc01::/58 dynamic network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: resCfgNetworkContainer_update2_ipv6,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_3",
					&ibclient.NetworkContainer{
						NetviewName: testNetView,
						Comment:     "dynamic network container testing",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config:      resCfgNetworkContainer_update3_ipv6,
				ExpectError: updateNotAllowedRegexp,
			},
		},
	})
}

func testAccCheckNetworkContainerDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ipv4_network_container" && rs.Type != "infoblox_ipv6_network_container" {
			continue
		}
		res, err := objMgr.GetNetworkContainerByRef(rs.Primary.ID)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if res != nil {
			return fmt.Errorf("object with ID '%s' remains", rs.Primary.ID)
		}
	}
	return nil
}
