package common

import (
	. "github.com/chivalryq/vela-go-sdk/pkg/apis"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApplicationBuilder struct {
	name        string
	namespace   string
	labels      map[string]string
	annotations map[string]string

	components   []Component
	steps        []WorkflowStep
	policies     []Policy
	workflowMode v1beta1.WorkflowExecuteMode
}

func (a *ApplicationBuilder) WithWorkflowSteps(step ...WorkflowStep) Application {
	a.steps = append(a.steps, step...)
	return a
}

// WithComponents append components to application
func (a *ApplicationBuilder) WithComponents(component ...Component) Application {
	a.components = append(a.components, component...)
	return a
}

// WithPolicies append policies to application
func (a *ApplicationBuilder) WithPolicies(policy ...Policy) Application {
	a.policies = append(a.policies, policy...)
	return a
}

// WithWorkflowMode set the workflow mode of application
func (a *ApplicationBuilder) WithWorkflowMode(steps, subSteps common.WorkflowMode) Application {
	a.workflowMode.Steps = steps
	a.workflowMode.SubSteps = subSteps
	return a
}

func (a *ApplicationBuilder) Name(name string) Application {
	a.name = name
	return a
}

func (a *ApplicationBuilder) Namespace(namespace string) Application {
	a.namespace = namespace
	return a
}

func (a *ApplicationBuilder) Labels(labels map[string]string) Application {
	a.labels = labels
	return a
}

func (a *ApplicationBuilder) Annotations(annotations map[string]string) Application {
	a.annotations = annotations
	return a
}

func (a *ApplicationBuilder) GetName() string {
	return a.name
}

func (a *ApplicationBuilder) GetNamespace() string {
	return a.namespace
}

func (a *ApplicationBuilder) GetLabels() map[string]string {
	return a.labels
}

func (a *ApplicationBuilder) GetAnnotations() map[string]string {
	return a.annotations
}

// New creates a new application with the given components.
func New() Application {
	app := &ApplicationBuilder{
		components: make([]Component, 0),
		steps:      make([]WorkflowStep, 0),
		policies:   make([]Policy, 0),
	}
	return app
}

func (a *ApplicationBuilder) Build() v1beta1.Application {
	components := make([]common.ApplicationComponent, 0, len(a.components))
	for _, component := range a.components {
		components = append(components, component.Build())
	}
	steps := make([]v1beta1.WorkflowStep, 0, len(a.steps))
	for _, step := range a.steps {
		steps = append(steps, step.Build())
	}
	policies := make([]v1beta1.AppPolicy, 0)
	for _, policy := range a.policies {
		policies = append(policies, policy.Build())
	}

	res := v1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       v1beta1.ApplicationKind,
			APIVersion: v1beta1.SchemeGroupVersion.String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      a.name,
			Namespace: a.namespace,
		},
		Spec: v1beta1.ApplicationSpec{
			Components: components,
			Workflow: &v1beta1.Workflow{
				Steps: steps,
			},
			Policies: policies,
		},
	}
	return res
}

func FromK8sObject(app *v1beta1.Application) (Application, error) {
	a := &ApplicationBuilder{}
	a.Name(app.Name)
	a.Namespace(app.Namespace)
	for _, comp := range app.Spec.Components {
		c, err := FromComponent(&comp)
		if err != nil {
			return nil, errors.Wrap(err, "convert component from k8s object")
		}
		a.WithComponents(c)
	}
	for _, step := range app.Spec.Workflow.Steps {
		s, err := FromWorkflowStep(&step)
		if err != nil {
			return nil, errors.Wrap(err, "convert workflow step from k8s object")
		}
		a.WithWorkflowSteps(s)
	}
	for _, policy := range app.Spec.Policies {
		p, err := FromPolicy(&policy)
		if err != nil {
			return nil, errors.Wrap(err, "convert policy from k8s object")
		}
		a.WithPolicies(p)
	}
	return a, nil
}

func FromComponent(component *common.ApplicationComponent) (Component, error) {
	build, ok := ComponentsBuilders[component.Type]
	if !ok {
		return nil, errors.Errorf("no component type %s registered", component.Type)
	}
	return build(*component)
}

func FromWorkflowStep(step *v1beta1.WorkflowStep) (WorkflowStep, error) {
	build, ok := WorkflowStepsBuilders[step.Type]
	if !ok {
		return nil, errors.Errorf("no workflow step type %s registered", step.Type)
	}
	return build(*step)
}

func FromPolicy(policy *v1beta1.AppPolicy) (Policy, error) {
	build, ok := PoliciesBuilders[policy.Type]
	if !ok {
		return nil, errors.Errorf("no policy type %s registered", policy.Type)
	}
	return build(*policy)
}

func FromTrait(trait *common.ApplicationTrait) (Trait, error) {
	build, ok := TraitBuilders[trait.Type]
	if !ok {
		return nil, errors.Errorf("no trait type %s registered", trait.Type)
	}
	return build(*trait)
}