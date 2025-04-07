/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/ajezierski/taskcli/internal/repository"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project management",
}

var createProjectCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := repository.CreateProject(args[0])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Println("Created project with ID:", id)
	},
}

var getProjectCmd = &cobra.Command{
	Use:   "get [project_id]",
	Short: "Show project details ",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := repository.GetProject(args[0])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if p == nil {
			fmt.Println("Project not found")
			return
		}
		fmt.Printf("ID: %s\nName: %s\n", p.ID, p.Name)
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update [project_id] [new_name]",
	Short: "Change project name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := repository.UpdateProject(args[0], args[1]); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Println("Project updated")
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "delete [project_id]",
	Short: "Delete project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := repository.DeleteProject(args[0]); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Println("Project deleted")
	},
}

var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all project",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := repository.ListProjects()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if len(projects) == 0 {
			fmt.Println("No projects\n")
			return
		}
		
		fmt.Printf("\n%-36s | %-30s\n", "PROJECT", "TITLE")
		fmt.Println(strings.Repeat("-", 78))
		
		for _, p := range projects {
			fmt.Printf("%-36s | %-30s\n", p.ID, p.Name)
		}
		fmt.Printf("\n")
	},
}

func init() {
	projectCmd.AddCommand(
		createProjectCmd,
		getProjectCmd,
		updateProjectCmd,
		deleteProjectCmd,
		listProjectsCmd,
	)

	rootCmd.AddCommand(projectCmd)
}
