package membership

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"den", "toss"}
		// NOTE: 위에서 이미 검증된 테스트케이스 이기 때문에 res, error 생략
		app.Create(CreateRequest{req.UserName, req.MembershipType})
		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})

		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "toss"}

		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})

		assert.NotNil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"den", ""}

		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})

		assert.NotNil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"den", "chai"}

		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})

		assert.NotNil(t, err)
	})
}

var defaultName = "den"
var defaultMembershipType = "toss"

func givenCreateDefaultNameMembership(app *Application) CreateResponse {
	req := CreateRequest{defaultName, defaultMembershipType}
	res, _ := app.Create(req)

	return res
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		var toUpdateName = "den_update"
		var toUpdateMembershipType = "naver"
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse := givenCreateDefaultNameMembership(app)

		req := UpdateRequest{createResponse.ID, toUpdateName, toUpdateMembershipType}
		res, err := app.Update(req)

		assert.Nil(t, err)
		assert.Equal(t, req.UserName, res.UserName)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		givenCreateDefaultNameMembership(app)
		createResponse, _ := app.Create(CreateRequest{"den2", "naver"})

		req := UpdateRequest{createResponse.ID, defaultName, "naver"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"", "den2", "naver"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"randomId", "", "toss"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"randomId", "den2", ""}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"randomId", "den2", "chai"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("업데이트를 하려는 멤버쉽 아이디가 없는 경우, 예외 처리 한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := UpdateRequest{"randomId", "den", "toss"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createResponse := givenCreateDefaultNameMembership(app)
		var lengthBeforeDelete = len(app.repository.data)

		err := app.Delete(createResponse.ID)

		assert.Nil(t, err)
		assert.Equal(t, lengthBeforeDelete-1, len(app.repository.data))
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		givenCreateDefaultNameMembership(app)

		err := app.Delete("")

		assert.NotNil(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		givenCreateDefaultNameMembership(app)

		err := app.Delete("randomId")

		assert.NotNil(t, err)
	})
}
