package main

import (
	"context"
	"fmt"
	"strings"

	todoist "github.com/sachaos/todoist/lib"
	"github.com/urfave/cli/v2"
)

func Modify(c *cli.Context) error {
	client := GetClient(c)

	if !c.Args().Present() {
		return CommandFailed
	}

	var err error
	item_id, err := client.CompleteItemIDByPrefix(c.Args().First())
	if err != nil {
		return err
	}
	item := client.Store.FindItem(item_id)
	if item == nil {
		return IdNotFound
	}
	if c.IsSet("content") {
		item.Content = c.String("content")
	}
	if c.IsSet("priority") {
		item.Priority = priorityMapping[c.Int("priority")]
	}
	if c.IsSet("label-names") {
		item.LabelNames = func(str string) []string {
			stringNames := strings.Split(str, ",")
			names := []string{}
			for _, stringName := range stringNames {
				if stringName != "" {
					names = append(names, stringName)
				}
			}
			return names
		}(c.String("label-names"))
	}

	if c.IsSet("date") {
		item.Due = &todoist.Due{String: c.String("date")}
	}

	// Resolve project ID
	projectID := c.String("project-id")
	if projectID == "" && c.IsSet("project-name") {
		projectID = client.Store.Projects.GetIDByName(c.String("project-name"))
		if projectID == "" {
			return fmt.Errorf("project not found: %s", c.String("project-name"))
		}
	}

	// Resolve section ID
	sectionID := c.String("section-id")
	if sectionID == "" && c.IsSet("section-name") {
		sectionName := c.String("section-name")
		// Find section by name within the target project (or current project if not moving)
		targetProjectID := projectID
		if targetProjectID == "" {
			targetProjectID = item.ProjectID
		}
		sectionID = client.Store.Sections.GetIDByNameAndProject(sectionName, targetProjectID)
		if sectionID == "" {
			return fmt.Errorf("section not found: %s", sectionName)
		}
	}

	// Update item properties
	if err := client.UpdateItem(context.Background(), *item); err != nil {
		return err
	}

	// Move to project if specified (must be done before section move)
	if projectID != "" {
		if err := client.MoveItem(context.Background(), item, projectID); err != nil {
			return err
		}
	}

	// Move to section if specified
	if sectionID != "" {
		if err := client.MoveItemToSection(context.Background(), item, sectionID); err != nil {
			return err
		}
	}

	return Sync(c)
}
