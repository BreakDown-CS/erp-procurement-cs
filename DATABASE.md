## suppliers

```
CREATE TABLE suppliers (
    id UUID PRIMARY KEY,
    supplier_code VARCHAR(50) UNIQUE,
    supplier_name VARCHAR(255),
    tax_id VARCHAR(50),
    email VARCHAR(255),
    phone VARCHAR(50),
    address TEXT,
    status VARCHAR(20) DEFAULT 'active',
    CONSTRAINT check_staff_status CHECK (
        status IN ('active', 'in_active')
    ),
    created_at TIMESTAMP DEFAULT NOW()
);
```

## purchase_requests

```
CREATE TABLE purchase_requests (
    id UUID PRIMARY KEY,
    pr_no VARCHAR(50) UNIQUE,
    request_date TIMESTAMP,
    requester_id UUID,
    department_id UUID,

    status VARCHAR(30) DEFAULT 'draft',
    CONSTRAINT check_staff_status CHECK (
        status IN ('draft', 'pending_approval', 'approved', 'rejected', 'cancelled', 'completed')
    ),

    remark TEXT,

    approved_by UUID,
    approved_at TIMESTAMP,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

## purchase_request_items

```
CREATE TABLE purchase_request_items (
    id UUID PRIMARY KEY,
    purchase_request_id UUID REFERENCES purchase_requests(id),

    item_name VARCHAR(255),
    qty NUMERIC(18,2),
    unit_price NUMERIC(18,2),

    total_price NUMERIC(18,2),

    created_at TIMESTAMP DEFAULT NOW()
);
```

## purchase_orders

```
CREATE TABLE purchase_orders (
    id UUID PRIMARY KEY,
    po_no VARCHAR(50) UNIQUE,

    purchase_request_id UUID REFERENCES purchase_requests(id),
    supplier_id UUID REFERENCES suppliers(id),

    order_date TIMESTAMP,

    status VARCHAR(30),

    total_amount NUMERIC(18,2),

    approved_by UUID,
    approved_at TIMESTAMP,

    created_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_permissions_code ON permissions(code);
```

## INDEX

## FK INDEXES
```
CREATE INDEX IF NOT EXISTS idx_purchase_request_items_request_id ON erp.purchase_request_items(purchase_request_id);
CREATE INDEX IF NOT EXISTS idx_purchase_orders_request_id ON erp.purchase_orders(purchase_request_id);
CREATE INDEX IF NOT EXISTS idx_purchase_orders_supplier_id ON erp.purchase_orders(supplier_id);
```

## STATUS INDEXES
```
CREATE INDEX IF NOT EXISTS idx_purchase_requests_status ON erp.purchase_requests(status);
CREATE INDEX IF NOT EXISTS idx_purchase_orders_status ON erp.purchase_orders(status);
```

## DATA INDEXES
```
CREATE INDEX IF NOT EXISTS idx_purchase_requests_request_date ON erp.purchase_requests(request_date);
CREATE INDEX IF NOT EXISTS idx_purchase_orders_order_date ON erp.purchase_orders(order_date);
```