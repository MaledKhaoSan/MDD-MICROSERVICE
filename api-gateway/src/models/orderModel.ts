export interface Order {
  customer_id: string;
  store_id: string;
  order_price: number;
  order_quantity: number;
  order_details: string;
}

export interface UpdateOrderStatus {
  status: string;
}

