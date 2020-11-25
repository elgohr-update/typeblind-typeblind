package github

import (
	"context"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"strings"
	"github.com/typeblind/typeblind/server/api_service/internal/consts"
	"github.com/typeblind/typeblind/server/api_service/internal/utils"
)

func GetRandomRepository(language string) (*github.Repository, error) {
	appliedOrgs := consts.AppliedOrgs()
	utils.ShuffleStrings(appliedOrgs)

	for i := range appliedOrgs {
		repo, err := tryToFindRepo(appliedOrgs[i], language)

		if err != nil {
			return nil, err
		} else if repo != nil {
			return repo, nil
		}
	}

	log.WithFields(log.Fields{
		"language": language,
	}).Info("Nothing was found for")
	return nil, nil
}


func tryToFindRepo(organization string, language string) (*github.Repository, error){
	ctx := context.Background()
	client := github.NewClient(nil)
	searchQuery := "org:" + organization + "+language:" + language
	log.Info(searchQuery)
	repos, _, err := client.Repositories.List(ctx, organization, nil)
	if err != nil {
		log.Error("RATE LIMIT ERROR")
		log.Error(err.Error())
		return nil, err
	}

	for i := range repos {
		lang := repos[i].GetLanguage()

		if strings.ToUpper(lang) == strings.ToUpper(language) {
			log.WithFields(log.Fields{
				"repo": repos[i],
			}).Info("Sucessfully find repo")
			return repos[i], err
		}
	}

	return nil, nil
}
