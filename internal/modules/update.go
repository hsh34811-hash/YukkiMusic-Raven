/*
 * ● YukkiMusic - Raven Edition
 * ○ High-performance Go-based Telegram music streaming bot.
 *
 * Modified by ✘ RAVEN
 * Repository: https://github.com/hsh34811-hash/YukkiMusic-Raven
 */

package modules

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Laky-64/gologging"
	tg "github.com/amarnathcjd/gogram/telegram"

	"main/internal/config"
	"main/internal/locales"
	"main/internal/utils"
)

func init() {
	helpTexts["/update"] = `<i>Check for updates and automatically upgrade the bot.</i>

<u>Usage:</u>
<b>/update</b> — Fetch latest commits, rebuild binary, and restart

<b>⚙️ Behavior:</b>
• Checks upstream GitHub repository for new commits
• Pulls latest code changes automatically
• Rebuilds Go application binary
• Hot-restarts the bot process without downtime

<b>🔒 Restrictions:</b>
• <b>Owner only</b> command`
	helpTexts["/gitpull"] = helpTexts["/update"]
}

func handleUpdate(m *tg.NewMessage) error {
	chatID := m.ChannelID()
	statusMsg, err := m.Reply(F(chatID, "update_checking"))
	if err != nil {
		gologging.Error("Failed to send update message: " + err.Error())
	}

	// 1. Ensure git is installed
	if _, err := exec.LookPath("git"); err != nil {
		utils.EOR(statusMsg, F(chatID, "update_no_git"))
		return tg.ErrEndGroup
	}

	upstream := config.UpstreamRepo
	branch := config.UpstreamBranch
	if upstream == "" {
		upstream = "https://github.com/hsh34811-hash/YukkiMusic-Raven.git"
	}
	if branch == "" {
		branch = "main"
	}

	// 2. Set upstream remote if not properly set
	_ = exec.Command("git", "remote", "set-url", "origin", upstream).Run()

	// 3. Fetch upstream branch
	fetchCmd := exec.Command("git", "fetch", "origin", branch)
	if out, err := fetchCmd.CombinedOutput(); err != nil {
		utils.EOR(statusMsg, F(chatID, "update_fetch_fail", locales.Arg{
			"error": string(out),
		}))
		return tg.ErrEndGroup
	}

	// 4. Count new commits
	revCountCmd := exec.Command("git", "rev-list", fmt.Sprintf("HEAD..origin/%s", branch), "--count")
	countBytes, _ := revCountCmd.Output()
	countStr := strings.TrimSpace(string(countBytes))

	if countStr == "" || countStr == "0" {
		utils.EOR(statusMsg, F(chatID, "update_uptodate"))
		return tg.ErrEndGroup
	}

	// 5. Get recent commit logs
	logCmd := exec.Command("git", "log", fmt.Sprintf("HEAD..origin/%s", branch), "--oneline", "-n", "5")
	logBytes, _ := logCmd.Output()
	commitsLog := strings.TrimSpace(string(logBytes))

	// Notify owner of found updates
	utils.EOR(statusMsg, F(chatID, "update_pulling", locales.Arg{
		"count": countStr,
		"log":   commitsLog,
	}))

	// 6. Stash local changes and pull
	_ = exec.Command("git", "stash").Run()
	pullCmd := exec.Command("git", "pull", "origin", branch)
	if out, err := pullCmd.CombinedOutput(); err != nil {
		utils.EOR(statusMsg, F(chatID, "update_pull_fail", locales.Arg{
			"error": string(out),
		}))
		return tg.ErrEndGroup
	}

	// 7. Check if Go compiler is available to rebuild binary
	if _, err := exec.LookPath("go"); err == nil {
		utils.EOR(statusMsg, F(chatID, "update_building"))
		buildCmd := exec.Command("go", "build", "-v", "-trimpath", "-ldflags=-w -s", "-o", "app", "./cmd/app/")
		if out, err := buildCmd.CombinedOutput(); err != nil {
			utils.EOR(statusMsg, F(chatID, "update_build_fail", locales.Arg{
				"error": string(out),
			}))
			return tg.ErrEndGroup
		}
	}

	// 8. Trigger graceful restart
	utils.EOR(statusMsg, F(chatID, "update_success", locales.Arg{
		"count": countStr,
	}))

	return executeRestart(m.Client, chatID, statusMsg)
}
