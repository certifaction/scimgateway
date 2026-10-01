package plugin

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/marcelom97/scimgateway/scim"
)

var testCtx = context.Background()

// contextAwarePlugin implements Plugin with context
type contextAwarePlugin struct {
	name string
}

func (p *contextAwarePlugin) Name() string { return p.name }

func (p *contextAwarePlugin) GetUsers(ctx context.Context, params scim.QueryParams) ([]*scim.User, error) {
	return []*scim.User{
		{ID: uuid.New().String(), UserName: "test"},
	}, nil
}

func (p *contextAwarePlugin) CreateUser(ctx context.Context, user *scim.User) (*scim.User, error) {
	user.ID = uuid.New().String()
	return user, nil
}

func (p *contextAwarePlugin) GetUser(ctx context.Context, id string, attributes []string) (*scim.User, error) {
	return &scim.User{ID: id, UserName: "test-user"}, nil
}

func (p *contextAwarePlugin) ModifyUser(ctx context.Context, id string, patch *scim.PatchOp) error {
	return nil
}

func (p *contextAwarePlugin) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (p *contextAwarePlugin) GetGroups(ctx context.Context, params scim.QueryParams) ([]*scim.Group, error) {
	return []*scim.Group{
		{ID: uuid.New().String(), DisplayName: "test-group"},
	}, nil
}

func (p *contextAwarePlugin) CreateGroup(ctx context.Context, group *scim.Group) (*scim.Group, error) {
	group.ID = uuid.New().String()
	return group, nil
}

func (p *contextAwarePlugin) GetGroup(ctx context.Context, id string, attributes []string) (*scim.Group, error) {
	return &scim.Group{ID: id, DisplayName: "test-group"}, nil
}

func (p *contextAwarePlugin) ModifyGroup(ctx context.Context, id string, patch *scim.PatchOp) error {
	return nil
}

func (p *contextAwarePlugin) DeleteGroup(ctx context.Context, id string) error {
	return nil
}

func TestNewAdapter(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	if adapter == nil {
		t.Fatal("NewAdapter() returned nil")
	}

	if adapter.plugin != p {
		t.Error("Adapter plugin not set correctly")
	}
}

func TestAdapterGetUsers(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	response, err := adapter.GetUsers(testCtx, scim.QueryParams{})
	if err != nil {
		t.Fatalf("GetUsers() error = %v", err)
	}

	if response.TotalResults != 1 {
		t.Errorf("Expected 1 result, got %d", response.TotalResults)
	}
}

func TestAdapterCreateUser(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	user := &scim.User{UserName: "newuser"}
	created, err := adapter.CreateUser(testCtx, user)
	if err != nil {
		t.Fatalf("CreateUser() error = %v", err)
	}

	if created.ID == "" {
		t.Error("Expected non-empty ID")
	}
}

func TestAdapterGetUser(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	testID := uuid.New().String()
	user, err := adapter.GetUser(testCtx, testID, nil)
	if err != nil {
		t.Fatalf("GetUser() error = %v", err)
	}

	if user.ID != testID {
		t.Errorf("Expected ID '%s', got '%s'", testID, user.ID)
	}

	if user.UserName != "test-user" {
		t.Errorf("Expected UserName 'test-user', got '%s'", user.UserName)
	}
}

func TestAdapterModifyUser(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	patch := &scim.PatchOp{
		Schemas: []string{scim.SchemaPatchOp},
		Operations: []scim.PatchOperation{
			{Op: "replace", Path: "active", Value: false},
		},
	}

	testID := uuid.New().String()
	err := adapter.ModifyUser(testCtx, testID, patch)
	if err != nil {
		t.Fatalf("ModifyUser() error = %v", err)
	}
}

func TestAdapterDeleteUser(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	testID := uuid.New().String()
	err := adapter.DeleteUser(testCtx, testID)
	if err != nil {
		t.Fatalf("DeleteUser() error = %v", err)
	}
}

func TestAdapterGetGroups(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	response, err := adapter.GetGroups(testCtx, scim.QueryParams{})
	if err != nil {
		t.Fatalf("GetGroups() error = %v", err)
	}

	if response.TotalResults != 1 {
		t.Errorf("Expected 1 result, got %d", response.TotalResults)
	}
}

func TestAdapterCreateGroup(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	group := &scim.Group{DisplayName: "newgroup"}
	created, err := adapter.CreateGroup(testCtx, group)
	if err != nil {
		t.Fatalf("CreateGroup() error = %v", err)
	}

	if created.ID == "" {
		t.Error("Expected non-empty ID")
	}
}

func TestAdapterGetGroup(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	testID := uuid.New().String()
	group, err := adapter.GetGroup(testCtx, testID, nil)
	if err != nil {
		t.Fatalf("GetGroup() error = %v", err)
	}

	if group.ID != testID {
		t.Errorf("Expected ID '%s', got '%s'", testID, group.ID)
	}

	if group.DisplayName != "test-group" {
		t.Errorf("Expected DisplayName 'test-group', got '%s'", group.DisplayName)
	}
}

func TestAdapterModifyGroup(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	patch := &scim.PatchOp{
		Schemas: []string{scim.SchemaPatchOp},
		Operations: []scim.PatchOperation{
			{Op: "replace", Path: "displayName", Value: "updated"},
		},
	}

	testID := uuid.New().String()
	err := adapter.ModifyGroup(testCtx, testID, patch)
	if err != nil {
		t.Fatalf("ModifyGroup() error = %v", err)
	}
}

func TestAdapterDeleteGroup(t *testing.T) {
	p := &contextAwarePlugin{name: "test"}
	adapter := NewAdapter(p)

	testID := uuid.New().String()
	err := adapter.DeleteGroup(testCtx, testID)
	if err != nil {
		t.Fatalf("DeleteGroup() error = %v", err)
	}
}

// trackingPlugin embeds contextAwarePlugin and counts delete/create calls so
// the fallback path can be asserted.
type trackingPlugin struct {
	contextAwarePlugin
	deleteUserCalls, createUserCalls   int
	deleteGroupCalls, createGroupCalls int
}

func (p *trackingPlugin) DeleteUser(ctx context.Context, id string) error {
	p.deleteUserCalls++
	return nil
}

func (p *trackingPlugin) CreateUser(ctx context.Context, user *scim.User) (*scim.User, error) {
	p.createUserCalls++
	return p.contextAwarePlugin.CreateUser(ctx, user)
}

func (p *trackingPlugin) DeleteGroup(ctx context.Context, id string) error {
	p.deleteGroupCalls++
	return nil
}

func (p *trackingPlugin) CreateGroup(ctx context.Context, group *scim.Group) (*scim.Group, error) {
	p.createGroupCalls++
	return p.contextAwarePlugin.CreateGroup(ctx, group)
}

// replacerPlugin adds the optional UserReplacer/GroupReplacer capabilities.
type replacerPlugin struct {
	trackingPlugin
	replaceUserCalls, replaceGroupCalls int
}

func (p *replacerPlugin) ReplaceUser(ctx context.Context, id string, user *scim.User) (*scim.User, error) {
	p.replaceUserCalls++
	user.ID = id
	return user, nil
}

func (p *replacerPlugin) ReplaceGroup(ctx context.Context, id string, group *scim.Group) (*scim.Group, error) {
	p.replaceGroupCalls++
	group.ID = id
	return group, nil
}

func TestAdapterReplaceUserDelegates(t *testing.T) {
	p := &replacerPlugin{}
	adapter := NewAdapter(p)

	replaced, err := adapter.ReplaceUser(testCtx, "u1", &scim.User{UserName: "updated"})
	if err != nil {
		t.Fatalf("ReplaceUser() error = %v", err)
	}
	if replaced.ID != "u1" {
		t.Errorf("Expected ID 'u1', got '%s'", replaced.ID)
	}
	if p.replaceUserCalls != 1 {
		t.Errorf("replaceUserCalls = %d, want 1", p.replaceUserCalls)
	}
	if p.deleteUserCalls != 0 || p.createUserCalls != 0 {
		t.Errorf("fallback used despite UserReplacer: delete=%d create=%d", p.deleteUserCalls, p.createUserCalls)
	}
}

func TestAdapterReplaceUserFallsBack(t *testing.T) {
	p := &trackingPlugin{}
	adapter := NewAdapter(p)

	replaced, err := adapter.ReplaceUser(testCtx, "u1", &scim.User{UserName: "updated"})
	if err != nil {
		t.Fatalf("ReplaceUser() error = %v", err)
	}
	if replaced == nil {
		t.Fatal("ReplaceUser() returned nil user")
	}
	if p.deleteUserCalls != 1 || p.createUserCalls != 1 {
		t.Errorf("legacy fallback must delete then create: delete=%d create=%d", p.deleteUserCalls, p.createUserCalls)
	}
}

func TestAdapterReplaceGroupDelegates(t *testing.T) {
	p := &replacerPlugin{}
	adapter := NewAdapter(p)

	replaced, err := adapter.ReplaceGroup(testCtx, "g1", &scim.Group{DisplayName: "updated"})
	if err != nil {
		t.Fatalf("ReplaceGroup() error = %v", err)
	}
	if replaced.ID != "g1" {
		t.Errorf("Expected ID 'g1', got '%s'", replaced.ID)
	}
	if p.replaceGroupCalls != 1 {
		t.Errorf("replaceGroupCalls = %d, want 1", p.replaceGroupCalls)
	}
	if p.deleteGroupCalls != 0 || p.createGroupCalls != 0 {
		t.Errorf("fallback used despite GroupReplacer: delete=%d create=%d", p.deleteGroupCalls, p.createGroupCalls)
	}
}

func TestAdapterReplaceGroupFallsBack(t *testing.T) {
	p := &trackingPlugin{}
	adapter := NewAdapter(p)

	replaced, err := adapter.ReplaceGroup(testCtx, "g1", &scim.Group{DisplayName: "updated"})
	if err != nil {
		t.Fatalf("ReplaceGroup() error = %v", err)
	}
	if replaced == nil {
		t.Fatal("ReplaceGroup() returned nil group")
	}
	if p.deleteGroupCalls != 1 || p.createGroupCalls != 1 {
		t.Errorf("legacy fallback must delete then create: delete=%d create=%d", p.deleteGroupCalls, p.createGroupCalls)
	}
}

func TestNewAdaptedManager(t *testing.T) {
	manager := NewManager()
	p := &contextAwarePlugin{name: "test"}
	manager.Register(p, nil)

	adaptedManager := NewAdaptedManager(manager)

	if adaptedManager == nil {
		t.Fatal("NewAdaptedManager() returned nil")
	}

	if adaptedManager.manager != manager {
		t.Error("AdaptedManager manager not set correctly")
	}
}

func TestAdaptedManagerGet(t *testing.T) {
	manager := NewManager()
	p := &contextAwarePlugin{name: "test"}
	manager.Register(p, nil)

	adaptedManager := NewAdaptedManager(manager)

	adapter, ok := adaptedManager.Get("test")
	if !ok {
		t.Fatal("Get() returned false for existing plugin")
	}

	if adapter == nil {
		t.Error("Get() returned nil adapter")
	}

	// Test that adapter works
	testID := uuid.New().String()
	user, err := adapter.GetUser(testCtx, testID, nil)
	if err != nil {
		t.Fatalf("Adapter GetUser() error = %v", err)
	}

	if user.ID != testID {
		t.Errorf("Expected ID '%s', got '%s'", testID, user.ID)
	}
}

func TestAdaptedManagerGetNonexistent(t *testing.T) {
	manager := NewManager()
	adaptedManager := NewAdaptedManager(manager)

	_, ok := adaptedManager.Get("nonexistent")
	if ok {
		t.Error("Get() returned true for non-existent plugin")
	}
}

func TestAdaptedManagerList(t *testing.T) {
	manager := NewManager()
	p1 := &contextAwarePlugin{name: "plugin1"}
	p2 := &contextAwarePlugin{name: "plugin2"}

	manager.Register(p1, nil)
	manager.Register(p2, nil)

	adaptedManager := NewAdaptedManager(manager)

	list := adaptedManager.List()
	if len(list) != 2 {
		t.Errorf("Expected 2 plugins, got %d", len(list))
	}
}
