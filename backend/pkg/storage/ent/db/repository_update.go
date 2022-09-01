// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/codecov"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowjobs"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/teams"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/workflows"
)

// RepositoryUpdate is the builder for updating Repository entities.
type RepositoryUpdate struct {
	config
	hooks    []Hook
	mutation *RepositoryMutation
}

// Where appends a list predicates to the RepositoryUpdate builder.
func (ru *RepositoryUpdate) Where(ps ...predicate.Repository) *RepositoryUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetRepositoryName sets the "repository_name" field.
func (ru *RepositoryUpdate) SetRepositoryName(s string) *RepositoryUpdate {
	ru.mutation.SetRepositoryName(s)
	return ru
}

// SetGitOrganization sets the "git_organization" field.
func (ru *RepositoryUpdate) SetGitOrganization(s string) *RepositoryUpdate {
	ru.mutation.SetGitOrganization(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepositoryUpdate) SetDescription(s string) *RepositoryUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetGitURL sets the "git_url" field.
func (ru *RepositoryUpdate) SetGitURL(s string) *RepositoryUpdate {
	ru.mutation.SetGitURL(s)
	return ru
}

// SetRepositoriesID sets the "repositories" edge to the Teams entity by ID.
func (ru *RepositoryUpdate) SetRepositoriesID(id uuid.UUID) *RepositoryUpdate {
	ru.mutation.SetRepositoriesID(id)
	return ru
}

// SetNillableRepositoriesID sets the "repositories" edge to the Teams entity by ID if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableRepositoriesID(id *uuid.UUID) *RepositoryUpdate {
	if id != nil {
		ru = ru.SetRepositoriesID(*id)
	}
	return ru
}

// SetRepositories sets the "repositories" edge to the Teams entity.
func (ru *RepositoryUpdate) SetRepositories(t *Teams) *RepositoryUpdate {
	return ru.SetRepositoriesID(t.ID)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflows entity by IDs.
func (ru *RepositoryUpdate) AddWorkflowIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddWorkflowIDs(ids...)
	return ru
}

// AddWorkflows adds the "workflows" edges to the Workflows entity.
func (ru *RepositoryUpdate) AddWorkflows(w ...*Workflows) *RepositoryUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.AddWorkflowIDs(ids...)
}

// AddCodecovIDs adds the "codecov" edge to the CodeCov entity by IDs.
func (ru *RepositoryUpdate) AddCodecovIDs(ids ...uuid.UUID) *RepositoryUpdate {
	ru.mutation.AddCodecovIDs(ids...)
	return ru
}

// AddCodecov adds the "codecov" edges to the CodeCov entity.
func (ru *RepositoryUpdate) AddCodecov(c ...*CodeCov) *RepositoryUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddCodecovIDs(ids...)
}

// AddProwSuiteIDs adds the "prow_suites" edge to the ProwSuites entity by IDs.
func (ru *RepositoryUpdate) AddProwSuiteIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddProwSuiteIDs(ids...)
	return ru
}

// AddProwSuites adds the "prow_suites" edges to the ProwSuites entity.
func (ru *RepositoryUpdate) AddProwSuites(p ...*ProwSuites) *RepositoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddProwSuiteIDs(ids...)
}

// AddProwJobIDs adds the "prow_jobs" edge to the ProwJobs entity by IDs.
func (ru *RepositoryUpdate) AddProwJobIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddProwJobIDs(ids...)
	return ru
}

// AddProwJobs adds the "prow_jobs" edges to the ProwJobs entity.
func (ru *RepositoryUpdate) AddProwJobs(p ...*ProwJobs) *RepositoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddProwJobIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ru *RepositoryUpdate) Mutation() *RepositoryMutation {
	return ru.mutation
}

// ClearRepositories clears the "repositories" edge to the Teams entity.
func (ru *RepositoryUpdate) ClearRepositories() *RepositoryUpdate {
	ru.mutation.ClearRepositories()
	return ru
}

// ClearWorkflows clears all "workflows" edges to the Workflows entity.
func (ru *RepositoryUpdate) ClearWorkflows() *RepositoryUpdate {
	ru.mutation.ClearWorkflows()
	return ru
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflows entities by IDs.
func (ru *RepositoryUpdate) RemoveWorkflowIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveWorkflowIDs(ids...)
	return ru
}

// RemoveWorkflows removes "workflows" edges to Workflows entities.
func (ru *RepositoryUpdate) RemoveWorkflows(w ...*Workflows) *RepositoryUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.RemoveWorkflowIDs(ids...)
}

// ClearCodecov clears all "codecov" edges to the CodeCov entity.
func (ru *RepositoryUpdate) ClearCodecov() *RepositoryUpdate {
	ru.mutation.ClearCodecov()
	return ru
}

// RemoveCodecovIDs removes the "codecov" edge to CodeCov entities by IDs.
func (ru *RepositoryUpdate) RemoveCodecovIDs(ids ...uuid.UUID) *RepositoryUpdate {
	ru.mutation.RemoveCodecovIDs(ids...)
	return ru
}

// RemoveCodecov removes "codecov" edges to CodeCov entities.
func (ru *RepositoryUpdate) RemoveCodecov(c ...*CodeCov) *RepositoryUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveCodecovIDs(ids...)
}

// ClearProwSuites clears all "prow_suites" edges to the ProwSuites entity.
func (ru *RepositoryUpdate) ClearProwSuites() *RepositoryUpdate {
	ru.mutation.ClearProwSuites()
	return ru
}

// RemoveProwSuiteIDs removes the "prow_suites" edge to ProwSuites entities by IDs.
func (ru *RepositoryUpdate) RemoveProwSuiteIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveProwSuiteIDs(ids...)
	return ru
}

// RemoveProwSuites removes "prow_suites" edges to ProwSuites entities.
func (ru *RepositoryUpdate) RemoveProwSuites(p ...*ProwSuites) *RepositoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemoveProwSuiteIDs(ids...)
}

// ClearProwJobs clears all "prow_jobs" edges to the ProwJobs entity.
func (ru *RepositoryUpdate) ClearProwJobs() *RepositoryUpdate {
	ru.mutation.ClearProwJobs()
	return ru
}

// RemoveProwJobIDs removes the "prow_jobs" edge to ProwJobs entities by IDs.
func (ru *RepositoryUpdate) RemoveProwJobIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveProwJobIDs(ids...)
	return ru
}

// RemoveProwJobs removes "prow_jobs" edges to ProwJobs entities.
func (ru *RepositoryUpdate) RemoveProwJobs(p ...*ProwJobs) *RepositoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemoveProwJobIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepositoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepositoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepositoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepositoryUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepositoryUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RepositoryUpdate) check() error {
	if v, ok := ru.mutation.RepositoryName(); ok {
		if err := repository.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf("db: validator failed for field \"repository_name\": %w", err)}
		}
	}
	if v, ok := ru.mutation.GitOrganization(); ok {
		if err := repository.GitOrganizationValidator(v); err != nil {
			return &ValidationError{Name: "git_organization", err: fmt.Errorf("db: validator failed for field \"git_organization\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Description(); ok {
		if err := repository.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("db: validator failed for field \"description\": %w", err)}
		}
	}
	if v, ok := ru.mutation.GitURL(); ok {
		if err := repository.GitURLValidator(v); err != nil {
			return &ValidationError{Name: "git_url", err: fmt.Errorf("db: validator failed for field \"git_url\": %w", err)}
		}
	}
	return nil
}

func (ru *RepositoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repository.Table,
			Columns: repository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repository.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RepositoryName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldRepositoryName,
		})
	}
	if value, ok := ru.mutation.GitOrganization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldGitOrganization,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldDescription,
		})
	}
	if value, ok := ru.mutation.GitURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldGitURL,
		})
	}
	if ru.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.RepositoriesTable,
			Columns: []string{repository.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teams.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.RepositoriesTable,
			Columns: []string{repository.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teams.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !ru.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedCodecovIDs(); len(nodes) > 0 && !ru.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedProwSuitesIDs(); len(nodes) > 0 && !ru.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProwSuitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ProwJobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedProwJobsIDs(); len(nodes) > 0 && !ru.mutation.ProwJobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProwJobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RepositoryUpdateOne is the builder for updating a single Repository entity.
type RepositoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepositoryMutation
}

// SetRepositoryName sets the "repository_name" field.
func (ruo *RepositoryUpdateOne) SetRepositoryName(s string) *RepositoryUpdateOne {
	ruo.mutation.SetRepositoryName(s)
	return ruo
}

// SetGitOrganization sets the "git_organization" field.
func (ruo *RepositoryUpdateOne) SetGitOrganization(s string) *RepositoryUpdateOne {
	ruo.mutation.SetGitOrganization(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepositoryUpdateOne) SetDescription(s string) *RepositoryUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetGitURL sets the "git_url" field.
func (ruo *RepositoryUpdateOne) SetGitURL(s string) *RepositoryUpdateOne {
	ruo.mutation.SetGitURL(s)
	return ruo
}

// SetRepositoriesID sets the "repositories" edge to the Teams entity by ID.
func (ruo *RepositoryUpdateOne) SetRepositoriesID(id uuid.UUID) *RepositoryUpdateOne {
	ruo.mutation.SetRepositoriesID(id)
	return ruo
}

// SetNillableRepositoriesID sets the "repositories" edge to the Teams entity by ID if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableRepositoriesID(id *uuid.UUID) *RepositoryUpdateOne {
	if id != nil {
		ruo = ruo.SetRepositoriesID(*id)
	}
	return ruo
}

// SetRepositories sets the "repositories" edge to the Teams entity.
func (ruo *RepositoryUpdateOne) SetRepositories(t *Teams) *RepositoryUpdateOne {
	return ruo.SetRepositoriesID(t.ID)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflows entity by IDs.
func (ruo *RepositoryUpdateOne) AddWorkflowIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddWorkflowIDs(ids...)
	return ruo
}

// AddWorkflows adds the "workflows" edges to the Workflows entity.
func (ruo *RepositoryUpdateOne) AddWorkflows(w ...*Workflows) *RepositoryUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.AddWorkflowIDs(ids...)
}

// AddCodecovIDs adds the "codecov" edge to the CodeCov entity by IDs.
func (ruo *RepositoryUpdateOne) AddCodecovIDs(ids ...uuid.UUID) *RepositoryUpdateOne {
	ruo.mutation.AddCodecovIDs(ids...)
	return ruo
}

// AddCodecov adds the "codecov" edges to the CodeCov entity.
func (ruo *RepositoryUpdateOne) AddCodecov(c ...*CodeCov) *RepositoryUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddCodecovIDs(ids...)
}

// AddProwSuiteIDs adds the "prow_suites" edge to the ProwSuites entity by IDs.
func (ruo *RepositoryUpdateOne) AddProwSuiteIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddProwSuiteIDs(ids...)
	return ruo
}

// AddProwSuites adds the "prow_suites" edges to the ProwSuites entity.
func (ruo *RepositoryUpdateOne) AddProwSuites(p ...*ProwSuites) *RepositoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddProwSuiteIDs(ids...)
}

// AddProwJobIDs adds the "prow_jobs" edge to the ProwJobs entity by IDs.
func (ruo *RepositoryUpdateOne) AddProwJobIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddProwJobIDs(ids...)
	return ruo
}

// AddProwJobs adds the "prow_jobs" edges to the ProwJobs entity.
func (ruo *RepositoryUpdateOne) AddProwJobs(p ...*ProwJobs) *RepositoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddProwJobIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ruo *RepositoryUpdateOne) Mutation() *RepositoryMutation {
	return ruo.mutation
}

// ClearRepositories clears the "repositories" edge to the Teams entity.
func (ruo *RepositoryUpdateOne) ClearRepositories() *RepositoryUpdateOne {
	ruo.mutation.ClearRepositories()
	return ruo
}

// ClearWorkflows clears all "workflows" edges to the Workflows entity.
func (ruo *RepositoryUpdateOne) ClearWorkflows() *RepositoryUpdateOne {
	ruo.mutation.ClearWorkflows()
	return ruo
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflows entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveWorkflowIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveWorkflowIDs(ids...)
	return ruo
}

// RemoveWorkflows removes "workflows" edges to Workflows entities.
func (ruo *RepositoryUpdateOne) RemoveWorkflows(w ...*Workflows) *RepositoryUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.RemoveWorkflowIDs(ids...)
}

// ClearCodecov clears all "codecov" edges to the CodeCov entity.
func (ruo *RepositoryUpdateOne) ClearCodecov() *RepositoryUpdateOne {
	ruo.mutation.ClearCodecov()
	return ruo
}

// RemoveCodecovIDs removes the "codecov" edge to CodeCov entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveCodecovIDs(ids ...uuid.UUID) *RepositoryUpdateOne {
	ruo.mutation.RemoveCodecovIDs(ids...)
	return ruo
}

// RemoveCodecov removes "codecov" edges to CodeCov entities.
func (ruo *RepositoryUpdateOne) RemoveCodecov(c ...*CodeCov) *RepositoryUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveCodecovIDs(ids...)
}

// ClearProwSuites clears all "prow_suites" edges to the ProwSuites entity.
func (ruo *RepositoryUpdateOne) ClearProwSuites() *RepositoryUpdateOne {
	ruo.mutation.ClearProwSuites()
	return ruo
}

// RemoveProwSuiteIDs removes the "prow_suites" edge to ProwSuites entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveProwSuiteIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveProwSuiteIDs(ids...)
	return ruo
}

// RemoveProwSuites removes "prow_suites" edges to ProwSuites entities.
func (ruo *RepositoryUpdateOne) RemoveProwSuites(p ...*ProwSuites) *RepositoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemoveProwSuiteIDs(ids...)
}

// ClearProwJobs clears all "prow_jobs" edges to the ProwJobs entity.
func (ruo *RepositoryUpdateOne) ClearProwJobs() *RepositoryUpdateOne {
	ruo.mutation.ClearProwJobs()
	return ruo
}

// RemoveProwJobIDs removes the "prow_jobs" edge to ProwJobs entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveProwJobIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveProwJobIDs(ids...)
	return ruo
}

// RemoveProwJobs removes "prow_jobs" edges to ProwJobs entities.
func (ruo *RepositoryUpdateOne) RemoveProwJobs(p ...*ProwJobs) *RepositoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemoveProwJobIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepositoryUpdateOne) Select(field string, fields ...string) *RepositoryUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repository entity.
func (ruo *RepositoryUpdateOne) Save(ctx context.Context) (*Repository, error) {
	var (
		err  error
		node *Repository
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepositoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) SaveX(ctx context.Context) *Repository {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepositoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RepositoryUpdateOne) check() error {
	if v, ok := ruo.mutation.RepositoryName(); ok {
		if err := repository.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf("db: validator failed for field \"repository_name\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.GitOrganization(); ok {
		if err := repository.GitOrganizationValidator(v); err != nil {
			return &ValidationError{Name: "git_organization", err: fmt.Errorf("db: validator failed for field \"git_organization\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Description(); ok {
		if err := repository.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("db: validator failed for field \"description\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.GitURL(); ok {
		if err := repository.GitURLValidator(v); err != nil {
			return &ValidationError{Name: "git_url", err: fmt.Errorf("db: validator failed for field \"git_url\": %w", err)}
		}
	}
	return nil
}

func (ruo *RepositoryUpdateOne) sqlSave(ctx context.Context) (_node *Repository, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repository.Table,
			Columns: repository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repository.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Repository.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for _, f := range fields {
			if !repository.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.RepositoryName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldRepositoryName,
		})
	}
	if value, ok := ruo.mutation.GitOrganization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldGitOrganization,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.GitURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldGitURL,
		})
	}
	if ruo.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.RepositoriesTable,
			Columns: []string{repository.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teams.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.RepositoriesTable,
			Columns: []string{repository.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teams.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !ruo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.WorkflowsTable,
			Columns: []string{repository.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflows.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedCodecovIDs(); len(nodes) > 0 && !ruo.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.CodecovTable,
			Columns: []string{repository.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: codecov.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedProwSuitesIDs(); len(nodes) > 0 && !ruo.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProwSuitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwSuitesTable,
			Columns: []string{repository.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowsuites.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ProwJobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedProwJobsIDs(); len(nodes) > 0 && !ruo.mutation.ProwJobsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProwJobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.ProwJobsTable,
			Columns: []string{repository.ProwJobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prowjobs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repository{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
