import axios from 'axios';
import { Order, UpdateOrderStatus } from '../models/orderModel';
import { Inventory } from '../models/inventoryModel';

const ORDER_EVENTS_SOURCING_URL = process.env.ORDER_EVENTS_SOURCING_URL;
const ORDER_HISTORY_SERVICE_URL = process.env.ORDER_HISTORY_SERVICE_URL;
const INVENTORY_SERVICE_URL = process.env.INVENTORY_SERVICE_URL;

if (!ORDER_HISTORY_SERVICE_URL) {
  throw new Error('ORDER_HISTORY_SERVICE_URL environment variable is not defined');
}

export const getOrderDetail = async (orderId: string): Promise<Order> => {
  console.log(`${ORDER_HISTORY_SERVICE_URL}/orders/${orderId}`);
  const response = await axios.get(`${ORDER_HISTORY_SERVICE_URL}/orders/${orderId}`);
  return response.data;
};

export const createOrder = async (orderData: Order): Promise<Order> => {
  const response = await axios.post(ORDER_HISTORY_SERVICE_URL, orderData);
  return response.data;
};

export const updateOrderStatus = async (orderId: string, statusData: UpdateOrderStatus): Promise<Order> => {
  const response = await axios.patch(`${ORDER_HISTORY_SERVICE_URL}/${orderId}/status`, statusData);
  return response.data;
};

export const AddItem = async (inventoryId: string, quantity: Inventory): Promise<Inventory> => {
  const response = await axios.patch(`${INVENTORY_SERVICE_URL}/${inventoryId}/inbound`, quantity);
  return response.data;
};

export const ReduceItem = async (inventoryId: string, quantity: Inventory): Promise<Inventory> => {
  const response = await axios.patch(`${INVENTORY_SERVICE_URL}/${inventoryId}/outbound`, quantity);
  return response.data;
};

