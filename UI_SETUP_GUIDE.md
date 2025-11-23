# 注册功能 UI 设置指南

## 需要在场景中添加的节点

在 `scenes/MainMenu.tscn` 的 `LoginPanel` 节点下，需要添加以下节点：

### 1. RegisterButton (Button)
- **位置**: LoginPanel 下，与 LoginButton 同级
- **位置坐标**: 可以放在 LoginButton 下方或旁边
- **脚本**: `res://scripts/ui/RegisterButton.gd`
- **文本**: "注册"
- **初始状态**: `visible = false`（默认隐藏，切换到注册模式时显示）

### 2. SwitchModeButton (Button)
- **位置**: LoginPanel 下
- **位置坐标**: 可以放在 LoginButton/RegisterButton 下方
- **文本**: "没有账号？注册"
- **功能**: 切换登录/注册模式

### 3. NameInput (LineEdit)
- **位置**: LoginPanel 下，在 PasswordInput 下方
- **位置坐标**: 
  - offset_left: 163.0
  - offset_top: 370.0 (在密码框下方)
  - offset_right: 401.0
  - offset_bottom: 401.0
- **初始状态**: `visible = false`（默认隐藏，切换到注册模式时显示）
- **提示文本**: "昵称（可选）"

### 4. NameLabel (Label)
- **位置**: LoginPanel 下，在 NameInput 左侧
- **位置坐标**:
  - offset_left: 95.0
  - offset_top: 369.0
  - offset_right: 155.0
  - offset_bottom: 402.0
- **文本**: "昵称"
- **初始状态**: `visible = false`

### 5. RegisterRequest (HTTPRequest)
- **位置**: LoginPanel 下，与 LoginRequest 同级
- **功能**: 处理注册请求

## 信号连接

需要在场景编辑器中连接以下信号：

1. **RegisterButton**:
   - `pressed` → `RegisterButton._on_pressed()`
   - `mouse_entered` → `RegisterButton._on_mouse_entered()`
   - `mouse_exited` → `RegisterButton._on_mouse_exited()`

2. **RegisterRequest**:
   - `request_completed` → `RegisterButton._on_request_completed()`

3. **SwitchModeButton**:
   - `pressed` → `LoginPanel._on_switch_mode_pressed()`

## 快速设置步骤

1. 打开 `scenes/MainMenu.tscn`
2. 在 `LoginPanel` 节点下添加上述节点
3. 设置节点的位置和属性
4. 连接信号
5. 保存场景

## 功能说明

- **登录模式**（默认）: 显示登录按钮，隐藏注册按钮和昵称输入框
- **注册模式**: 显示注册按钮和昵称输入框，隐藏登录按钮
- 点击"没有账号？注册"按钮可以切换模式
- 注册成功后会自动登录并关闭登录面板

