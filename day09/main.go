package main

import (
    "fmt"
)

type ICoupons interface {
    ApplyCoupons() float64
}

// BaseOrder 是订单的基类，实现了模板方法
type BaseOrder struct {
    Items         []Item
    DiscountCodes []string // 未使用，仅为示例
    TaxRate       float64
    Coupons       ICoupons // 改为引用 ICoupons 接口
}

// Item 是订单中的商品项
type Item struct {
    Price    float64
    Quantity int
}

// CalculateTotal 计算所有商品的总价
func (o *BaseOrder) CalculateTotal() float64 {
    total := 0.0
    for _, item := range o.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

// CalculateTaxes 计算税费
func (o *BaseOrder) CalculateTaxes(totalAfterDiscount float64) float64 {
    return totalAfterDiscount * o.TaxRate
}

// GetFinalAmount 模板方法，定义了订单结算的流程
func (o *BaseOrder) GetFinalAmount() float64 {
    total := o.CalculateTotal()
    discount := o.Coupons.ApplyCoupons()
    fmt.Println(total)
    fmt.Println(discount)
    totalAfterDiscount := total - discount
    taxes := o.CalculateTaxes(totalAfterDiscount)
    finalAmount := totalAfterDiscount + taxes
    return finalAmount
}

// VIPOrder 是VIP用户的订单，它可以重写优惠券应用方法
type VIPOrder struct {
    *BaseOrder // 嵌套 BaseOrder
}

// ApplyCoupons 为VIP用户应用特殊优惠券
func (o *VIPOrder) ApplyCoupons() float64 {
    discount := o.CalculateTotal() * 0.10
    return discount
}

type NormalOrder struct {
    *BaseOrder // 嵌套 BaseOrder
}

func (o *NormalOrder) ApplyCoupons() float64 {
    // 普通订单不应用任何优惠
    return 0
}

func main() {
    // 创建一个普通订单
    items := []Item{{Price: 100, Quantity: 2}, {Price: 200, Quantity: 1}}
    normalOrder := &NormalOrder{BaseOrder: &BaseOrder{Items: items, TaxRate: 0.05}}
    normalOrder.Coupons = normalOrder // 将普通订单赋值给 Coupons 字段
    fmt.Printf("Normal Order Final Amount: %.2f\n", normalOrder.GetFinalAmount())

    // 创建一个VIP订单
    vipOrder := &VIPOrder{BaseOrder: &BaseOrder{Items: items, TaxRate: 0.05}}
    vipOrder.Coupons = vipOrder // 将VIP订单赋值给 Coupons 字段
    fmt.Printf("VIP Order Final Amount: %.2f\n", vipOrder.GetFinalAmount())
}
