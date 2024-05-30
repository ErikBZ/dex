// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/devicerequest"
)

// DeviceRequestCreate is the builder for creating a DeviceRequest entity.
type DeviceRequestCreate struct {
	config
	mutation *DeviceRequestMutation
	hooks    []Hook
}

// SetUserCode sets the "user_code" field.
func (drc *DeviceRequestCreate) SetUserCode(s string) *DeviceRequestCreate {
	drc.mutation.SetUserCode(s)
	return drc
}

// SetDeviceCode sets the "device_code" field.
func (drc *DeviceRequestCreate) SetDeviceCode(s string) *DeviceRequestCreate {
	drc.mutation.SetDeviceCode(s)
	return drc
}

// SetClientID sets the "client_id" field.
func (drc *DeviceRequestCreate) SetClientID(s string) *DeviceRequestCreate {
	drc.mutation.SetClientID(s)
	return drc
}

// SetClientSecret sets the "client_secret" field.
func (drc *DeviceRequestCreate) SetClientSecret(s string) *DeviceRequestCreate {
	drc.mutation.SetClientSecret(s)
	return drc
}

// SetScopes sets the "scopes" field.
func (drc *DeviceRequestCreate) SetScopes(s []string) *DeviceRequestCreate {
	drc.mutation.SetScopes(s)
	return drc
}

// SetExpiry sets the "expiry" field.
func (drc *DeviceRequestCreate) SetExpiry(t time.Time) *DeviceRequestCreate {
	drc.mutation.SetExpiry(t)
	return drc
}

// Mutation returns the DeviceRequestMutation object of the builder.
func (drc *DeviceRequestCreate) Mutation() *DeviceRequestMutation {
	return drc.mutation
}

// Save creates the DeviceRequest in the database.
func (drc *DeviceRequestCreate) Save(ctx context.Context) (*DeviceRequest, error) {
	return withHooks(ctx, drc.sqlSave, drc.mutation, drc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (drc *DeviceRequestCreate) SaveX(ctx context.Context) *DeviceRequest {
	v, err := drc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drc *DeviceRequestCreate) Exec(ctx context.Context) error {
	_, err := drc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drc *DeviceRequestCreate) ExecX(ctx context.Context) {
	if err := drc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drc *DeviceRequestCreate) check() error {
	if _, ok := drc.mutation.UserCode(); !ok {
		return &ValidationError{Name: "user_code", err: errors.New(`db: missing required field "DeviceRequest.user_code"`)}
	}
	if v, ok := drc.mutation.UserCode(); ok {
		if err := devicerequest.UserCodeValidator(v); err != nil {
			return &ValidationError{Name: "user_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.user_code": %w`, err)}
		}
	}
	if _, ok := drc.mutation.DeviceCode(); !ok {
		return &ValidationError{Name: "device_code", err: errors.New(`db: missing required field "DeviceRequest.device_code"`)}
	}
	if v, ok := drc.mutation.DeviceCode(); ok {
		if err := devicerequest.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.device_code": %w`, err)}
		}
	}
	if _, ok := drc.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`db: missing required field "DeviceRequest.client_id"`)}
	}
	if v, ok := drc.mutation.ClientID(); ok {
		if err := devicerequest.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_id": %w`, err)}
		}
	}
	if _, ok := drc.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`db: missing required field "DeviceRequest.client_secret"`)}
	}
	if v, ok := drc.mutation.ClientSecret(); ok {
		if err := devicerequest.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_secret": %w`, err)}
		}
	}
	if _, ok := drc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`db: missing required field "DeviceRequest.expiry"`)}
	}
	return nil
}

func (drc *DeviceRequestCreate) sqlSave(ctx context.Context) (*DeviceRequest, error) {
	if err := drc.check(); err != nil {
		return nil, err
	}
	_node, _spec := drc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	drc.mutation.id = &_node.ID
	drc.mutation.done = true
	return _node, nil
}

func (drc *DeviceRequestCreate) createSpec() (*DeviceRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceRequest{config: drc.config}
		_spec = sqlgraph.NewCreateSpec(devicerequest.Table, sqlgraph.NewFieldSpec(devicerequest.FieldID, field.TypeInt))
	)
	if value, ok := drc.mutation.UserCode(); ok {
		_spec.SetField(devicerequest.FieldUserCode, field.TypeString, value)
		_node.UserCode = value
	}
	if value, ok := drc.mutation.DeviceCode(); ok {
		_spec.SetField(devicerequest.FieldDeviceCode, field.TypeString, value)
		_node.DeviceCode = value
	}
	if value, ok := drc.mutation.ClientID(); ok {
		_spec.SetField(devicerequest.FieldClientID, field.TypeString, value)
		_node.ClientID = value
	}
	if value, ok := drc.mutation.ClientSecret(); ok {
		_spec.SetField(devicerequest.FieldClientSecret, field.TypeString, value)
		_node.ClientSecret = value
	}
	if value, ok := drc.mutation.Scopes(); ok {
		_spec.SetField(devicerequest.FieldScopes, field.TypeJSON, value)
		_node.Scopes = value
	}
	if value, ok := drc.mutation.Expiry(); ok {
		_spec.SetField(devicerequest.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	return _node, _spec
}

// DeviceRequestCreateBulk is the builder for creating many DeviceRequest entities in bulk.
type DeviceRequestCreateBulk struct {
	config
	builders []*DeviceRequestCreate
}

// Save creates the DeviceRequest entities in the database.
func (drcb *DeviceRequestCreateBulk) Save(ctx context.Context) ([]*DeviceRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(drcb.builders))
	nodes := make([]*DeviceRequest, len(drcb.builders))
	mutators := make([]Mutator, len(drcb.builders))
	for i := range drcb.builders {
		func(i int, root context.Context) {
			builder := drcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, drcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, drcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, drcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (drcb *DeviceRequestCreateBulk) SaveX(ctx context.Context) []*DeviceRequest {
	v, err := drcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drcb *DeviceRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := drcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drcb *DeviceRequestCreateBulk) ExecX(ctx context.Context) {
	if err := drcb.Exec(ctx); err != nil {
		panic(err)
	}
}