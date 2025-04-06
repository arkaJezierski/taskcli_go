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

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var creareTaskCmd = &cobra.Command{
	Use:   "create [project_id] [title]",
	Short: "Create a task in a project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := repository.CreateTaskToProject(args[0], args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task created with ID:", id)
	},
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		projectID, _ := flags.GetString("project")

		tasks, err := repository.ListTasks(projectID)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Printf("\n%-36s  %-30s %-12s  %-6s\n", "TASK ID", "TITLE", "PROJECT", "STATUS")
		// Reapest for every param + six for spaces
		fmt.Println(strings.Repeat("-", 36+30+12+6+6))

		for _, t := range tasks {
			// coloring by using git shell coloring 
			status := "\033[31mTODO\033[0m"
			if t.Done {
				status = "\033[32mDONE\033[0m"
			}
			fmt.Printf("%-36s  %-30s %-12s %-6s\n", t.ID, t.Title, t.ProjectID, status)
		}
		fmt.Println()
	},
}

var doneTaskCmd = &cobra.Command{
	Use:   "done [task_id]",
	Short: "Mark task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := repository.MarkTaskDone(args[0]); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task marked as done.")
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "Delete task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := repository.DeleteTask(args[0]); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task deleted.")
	},
}

func init() {
	listTasksCmd.Flags().String("project", "", "Filter by project ID")
	taskCmd.AddCommand(creareTaskCmd, listTasksCmd, doneTaskCmd, deleteTaskCmd)
	rootCmd.AddCommand(taskCmd)
}
