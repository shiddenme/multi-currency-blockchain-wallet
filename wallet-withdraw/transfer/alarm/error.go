package alarm

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// ErrFeeNotEnough
type ErrorBalanceNotEnough struct {
	ErrorDetail string
}

func NewErrorBalanceNotEnough(balance, amount interface{}) *ErrorBalanceNotEnough {
	return &ErrorBalanceNotEnough{
		ErrorDetail: fmt.Sprintf("系统地址余额不足,当前余额为:[%v],交易金额为:[%v]", balance, amount),
	}
}

func (e *ErrorBalanceNotEnough) Error() string {
	return fmt.Sprintf("wallet balance not enough")
}

// ErrFeeNotEnough def.
type ErrorBalanceLessThanFee struct {
	ErrorDetail string
}

func NewErrorBalanceLessThanFee(fee decimal.Decimal) *ErrorBalanceLessThanFee {
	return &ErrorBalanceLessThanFee{
		ErrorDetail: fmt.Sprintf("无可用的系统地址: 系统地址资金小于最低交易手续费;当前交易手续费为:[%v]", fee),
	}
}

func (e *ErrorBalanceLessThanFee) Error() string {
	return fmt.Sprintf("no avaliable accounts for system ")
}

// ErrorBalanceLessCost.
type ErrorBalanceLessCost struct {
	ErrorDetail string
}

func NewErrorBalanceLessCost(fee, balance decimal.Decimal) *ErrorBalanceLessCost {
	return &ErrorBalanceLessCost{
		ErrorDetail: fmt.Sprintf("系统地址余额不足: 系统地址余额无法支付交易所需手续费, 当前所需手续费为 %v ;满足账户的资金合计为 %v", fee, balance),
	}
}

func (e *ErrorBalanceLessCost) Error() string {
	return fmt.Sprintf("wallet balance less than cost ")
}

// ErrorBalanceLessCost.
type NotMatchAccount struct {
	ErrorDetail string
}

func NewNotMatchAccount(fee, amount, balance decimal.Decimal, address string) *NotMatchAccount {
	difference := amount.Sub(balance)
	return &NotMatchAccount{
		ErrorDetail: fmt.Sprintf("未匹配到满足当前交易的系统地址; 当前交易金额为: %v ;手续费为: %v ;当前系统地址中最大金额为: %v,地址为:[%s],至少需要再转入金额为:[%v]", fee, balance, amount, address, difference),
	}
}

func (e *NotMatchAccount) Error() string {
	return fmt.Sprintf("don't match the appropriate system address ")
}

type ErrorTxFeeNotEnough struct {
	ErrorDetail string
}

func NewErrorTxFeeNotEnough(fee, needFee decimal.Decimal) *ErrorTxFeeNotEnough {
	return &ErrorTxFeeNotEnough{
		ErrorDetail: fmt.Sprintf("交易手续费不足; 当前手续费为: %v; 需要手续费为: %v ", fee, needFee),
	}
}

func (e *ErrorTxFeeNotEnough) Error() string {
	return fmt.Sprintf("tx transfer fee not enough ")
}

type ErrorAccountBalanceNotEnough struct {
	ErrorDetail string
}

func NewErrorAccountBalanceNotEnough(address string, balance, needFee decimal.Decimal) *ErrorAccountBalanceNotEnough {
	return &ErrorAccountBalanceNotEnough{
		ErrorDetail: fmt.Sprintf("交易失败,地址余额低于当前手续费; 地址: %s ,余额 %v ,当前交易手续费: %v ", address, balance, needFee),
	}
}

func (e *ErrorAccountBalanceNotEnough) Error() string {
	return fmt.Sprintf("address balance not enough for transaction fee ")
}

type WarnCoolWalletBalanceChange struct {
	WarnDetail string
}

func NewWarnCoolWalletBalanceChange(maxRemainBalance, amount decimal.Decimal, coldAddress, c string) *WarnCoolWalletBalanceChange {
	return &WarnCoolWalletBalanceChange{
		WarnDetail: fmt.Sprintf("钱包资金超过预留最大值: %s, 当前余额为:%s, 已转入冷钱包(%s),交易哈希值为:%s", maxRemainBalance.String(), amount.String(), coldAddress, c),
	}
}

func (e *WarnCoolWalletBalanceChange) Error() string {
	return fmt.Sprintf("cool wallet balance has changed ")
}
