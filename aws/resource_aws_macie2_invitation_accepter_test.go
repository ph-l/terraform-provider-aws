package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccAwsMacie2InvitationAccepter_basic(t *testing.T) {
	var providers []*schema.Provider
	resourceName := "aws_macie2_invitation_accepter.test"
	email := "required@example.com"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAlternateAccountPreCheck(t)
		},
		ProviderFactories: testAccProviderFactoriesAlternate(&providers),
		CheckDestroy:      testAccCheckAwsMacie2InvitationAccepterDestroy,
		ErrorCheck:        testAccErrorCheck(t, macie2.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsMacieInvitationAccepterConfigBasic(email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMacie2InvitationAccepterExists(resourceName),
				),
			},
			{
				Config:            testAccAwsMacieInvitationAccepterConfigBasic(email),
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAwsMacie2InvitationAccepter_memberStatus(t *testing.T) {
	var providers []*schema.Provider
	resourceName := "aws_macie2_invitation_accepter.test"
	email := "required@example.com"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAlternateAccountPreCheck(t)
		},
		ProviderFactories: testAccProviderFactoriesAlternate(&providers),
		CheckDestroy:      testAccCheckAwsMacie2InvitationAccepterDestroy,
		ErrorCheck:        testAccErrorCheck(t, macie2.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsMacieInvitationAccepterConfigMemberStatus(email, macie2.MacieStatusEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMacie2InvitationAccepterExists(resourceName),
				),
			},
			{
				Config: testAccAwsMacieInvitationAccepterConfigMemberStatus(email, macie2.MacieStatusPaused),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMacie2InvitationAccepterExists(resourceName),
				),
			},
			{
				Config:            testAccAwsMacieInvitationAccepterConfigBasic(email),
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAwsMacie2InvitationAccepterExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource (%s) has empty ID", resourceName)
		}

		conn := testAccProvider.Meta().(*AWSClient).macie2conn
		input := &macie2.GetAdministratorAccountInput{}
		output, err := conn.GetAdministratorAccount(input)

		if err != nil {
			return err
		}

		if output == nil || output.Administrator == nil || aws.StringValue(output.Administrator.AccountId) == "" {
			return fmt.Errorf("no administrator account found for: %s", resourceName)
		}

		return nil
	}
}

func testAccCheckAwsMacie2InvitationAccepterDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).macie2conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_macie2_invitation_accepter" {
			continue
		}

		input := &macie2.GetAdministratorAccountInput{}
		output, err := conn.GetAdministratorAccount(input)
		if tfawserr.ErrCodeEquals(err, macie2.ErrCodeResourceNotFoundException) ||
			tfawserr.ErrMessageContains(err, macie2.ErrCodeAccessDeniedException, "Macie is not enabled") {
			continue
		}

		if output == nil || output.Administrator == nil || aws.StringValue(output.Administrator.AccountId) != rs.Primary.Attributes["administrator_account_id"] {
			continue
		}

		return fmt.Errorf("macie InvitationAccepter %q still exists", rs.Primary.ID)

	}

	return nil

}

func testAccAwsMacieInvitationAccepterConfigBasic(email string) string {
	return testAccAlternateAccountProviderConfig() + fmt.Sprintf(`
data "aws_caller_identity" "primary" {
  provider = "awsalternate"
}

data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "primary" {
  provider = "awsalternate"
}

resource "aws_macie2_account" "member" {}

resource "aws_macie2_member" "primary" {
  provider   = "awsalternate"
  account_id = data.aws_caller_identity.current.account_id
  email      = %[1]q
  depends_on = [aws_macie2_account.primary]
}

resource "aws_macie2_invitation" "primary" {
  provider   = "awsalternate"
  account_id = data.aws_caller_identity.current.account_id
  depends_on = [aws_macie2_member.primary]
}

resource "aws_macie2_invitation_accepter" "test" {
  administrator_account_id = data.aws_caller_identity.primary.account_id
  depends_on               = [aws_macie2_invitation.primary]
}

`, email)
}

func testAccAwsMacieInvitationAccepterConfigMemberStatus(email, memberStatus string) string {
	return testAccAlternateAccountProviderConfig() + fmt.Sprintf(`
data "aws_caller_identity" "primary" {
  provider = "awsalternate"
}

data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "primary" {
  provider = "awsalternate"
}

resource "aws_macie2_account" "member" {}

resource "aws_macie2_member" "primary" {
  provider   = "awsalternate"
  account_id = data.aws_caller_identity.current.account_id
  email      = %[1]q
  status     = %[2]q
  depends_on = [aws_macie2_account.primary]
}

resource "aws_macie2_invitation" "primary" {
  provider   = "awsalternate"
  account_id = data.aws_caller_identity.current.account_id
  depends_on = [aws_macie2_member.primary]
}

resource "aws_macie2_invitation_accepter" "test" {
  administrator_account_id = data.aws_caller_identity.primary.account_id
  depends_on               = [aws_macie2_invitation.primary]
}

`, email, memberStatus)
}
