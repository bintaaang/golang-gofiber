package models

import (
    "time"
    "gorm.io/gorm"
)

// User untuk admin dan kurir
type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    Name      string         `json:"name"`
    Email     string         `json:"email"`
    Phone     string         `json:"phone"`
    Role      string         `json:"role"` 
    Password  string         `json:"password"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

// PickupRequest untuk user luar
type PickupRequest struct {
    ID           uint           `json:"id" gorm:"primaryKey"`
    TrackingNo   string         `json:"tracking_no" gorm:"uniqueIndex"`
    Name         string         `json:"name"`        // Nama pengirim
    Phone        string         `json:"phone"`       // No HP pengirim
    AddressFrom  string         `json:"address_from"`
    AddressTo    string         `json:"address_to"`
    CourierID    *uint          `json:"courier_id"`  // nullable sampai diassign
    Courier      *User          `gorm:"foreignKey:CourierID"`
    Status       string         `json:"status"`      // pending, assigned, picked_up, in_sorting, in_transit, on_delivery, delivered, return, undelivered
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
    Mark         uint           `json:"mark"`
}

// PackageStatus untuk log perubahan status paket
type PackageStatus struct {
    ID             uint           `json:"id" gorm:"primaryKey"`
    PickupRequestID uint          `json:"pickup_request_id"`
    TrackingNo     string         `json:"tracking_no" gorm:"index"` // copy dari PickupRequest
    PickupRequest  PickupRequest  `gorm:"foreignKey:PickupRequestID"`
    Status         string         `json:"status"`         // Sama seperti PickupRequest.Status
    UpdatedByID    uint           `json:"updated_by_id"`  // User/admin/kurir ID
    UpdatedBy      User           `gorm:"foreignKey:UpdatedByID"`
    Note           string         `json:"note"`           // Catatan optional
    CreatedAt      time.Time
}

// Enum status opsional
var PickupStatuses = struct {
    Pending      string
    Assigned     string
    PickedUp     string
    InSorting    string
    InTransit    string
    OnDelivery   string
    Delivered    string
    Return       string
    Undelivered  string
}{
    Pending:     "pending",
    Assigned:    "assigned", //kurir
    PickedUp:    "picked_up",
    InSorting:   "in_sorting",
    InTransit:   "in_transit",
    OnDelivery:  "on_delivery", //kurir
    Delivered:   "delivered", ///kuri
    Return:      "return",
    Undelivered: "undelivered",
}
