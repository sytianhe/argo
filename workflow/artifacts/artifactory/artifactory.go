package artifactory

import (
	wfv1 "github.com/argoproj/argo/api/workflow/v1"

	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/services/utils/auth"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/services/utils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/services"
	"github.com/argoproj/argo/errors"
)


// ArtifactoryArtifactDriver is a driver for JFrog Artifactory
type ArtifactoryArtifactDriver struct {
	Endpoint string
	User     string
	Password string
}

func (artifactoryDriver *ArtifactoryArtifactDriver) newArtifactoryClient() (*artifactory.ArtifactoryServicesManager, error) {

	rtDetails := new(auth.ArtifactoryDetails)
	rtDetails.Url = artifactoryDriver.Endpoint
	rtDetails.User = artifactoryDriver.User
	rtDetails.Password = artifactoryDriver.Password

	serviceConfig, err := (&artifactory.ArtifactoryServicesConfigBuilder{}).
		SetArtDetails(rtDetails).
		SetDryRun(false).
		Build()
	if err != nil {
		return nil, errors.InternalWrapError(err, "Failed to create Artifactory config builder.")
	}

	rtManager, err := artifactory.NewArtifactoryService(serviceConfig)
	if err != nil {
		return nil, errors.InternalWrapError(err, "Failed to create Artifactory client manager.")
	}
	return rtManager, nil
}

func (artifactoryDriver *ArtifactoryArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	rtManager, err := artifactoryDriver.newArtifactoryClient()
	if err != nil {
		return err
	}
	params := new(utils.ArtifactoryCommonParams)
	params.Pattern = inputArtifact.Artifactory.RepoPattern
	params.Target = path
	downloadParams := &services.DownloadParamsImpl{}
	downloadParams.ArtifactoryCommonParams = params
	rtManager.DownloadFiles(downloadParams)
	return nil
}

func (artifactoryDriver *ArtifactoryArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	rtManager, err := artifactoryDriver.newArtifactoryClient()
	if err != nil {
		return err
	}
	params := new(utils.ArtifactoryCommonParams)
	params.Pattern = outputArtifact.Artifactory.RepoPattern
	params.Target = path
	uploadParam := &services.UploadParamsImp{}
	uploadParam.ArtifactoryCommonParams = params
	rtManager.UploadFiles(uploadParam)
	return nil
}
