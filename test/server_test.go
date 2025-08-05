package main

import (
	"os/exec"
	"testing"
	"time"
)

func TestChatboxTool(t *testing.T) {
	cmd := exec.Command("make", "mcphost")

	// 可选：打印输出，便于调试
	cmd.Stdout = nil // 可以换成 os.Stdout 看启动日志
	cmd.Stderr = nil

	// 启动命令（非阻塞）
	if err := cmd.Start(); err != nil {
		t.Fatalf("❌ 启动 mcphost 失败: %v", err)
	}

	// 等待几秒，确保进程已启动（可根据启动速度调整）
	time.Sleep(10 * time.Second)
	t.Logf("✅ mcphost 启动成功，进程ID: %d", cmd.Process.Pid)

	// 可选：测试完是否结束进程
	defer func() {
		_ = cmd.Process.Kill()
		t.Log("🛑 已终止 mcphost 进程")
	}()
}
