import { Request, Response } from 'express';

import { getOrderDetail, createOrder, updateOrderStatus } from '../services/apiService';

import { Order, UpdateOrderStatus } from '../models/orderModel';

export const getOrderById = async (req: Request, res: Response) => {
  try {
    const orderId = req.params.id;
    const order = await getOrderDetail(orderId);
    res.status(200).json(order);
  } catch (error) {
    res.status(500).send('Error fetching order');
  }
};

export const postOrder = async (req: Request, res: Response) => {
  try {
    const orderData: Order = req.body;

  if (!orderData.customer_id || typeof orderData.customer_id !== 'string') {
    res.status(400).json({ error: 'Customer ID is required and must be a string' });
    return;
  }

  if (!orderData.store_id || typeof orderData.store_id !== 'string') {
    res.status(400).json({ error: 'Store ID is required and must be a string' });
    return;
  }

  if (typeof orderData.order_price !== 'number' || orderData.order_price <= 0) {
    res.status(400).json({ error: 'Order price must be a positive number' });
    return;
  }

  if (!Number.isInteger(orderData.order_quantity) || orderData.order_quantity <= 0) {
    res.status(400).json({ error: 'Order quantity must be a positive integer' });
    return;
  }

  if (!orderData.order_details || typeof orderData.order_details !== 'string' || orderData.order_details.length < 5) {
    res.status(400).json({ error: 'Order details must be at least 5 characters long' });
    return;
  }

    const newOrder = await createOrder(orderData);
    res.status(201).json({ message: 'Order created successfully', order: newOrder });
  } catch (error) {
    res.status(500).send('Error creating order');
  }
};

export const patchOrderStatus = async (req: Request, res: Response): Promise<void> => {
  try {
    const orderId = req.params.id;
    const statusData: UpdateOrderStatus = req.body;

    if (!orderId) {
      res.status(400).json({ error: 'Invalid order ID' });
      return;
    }

    if (!statusData.status) {
      res.status(400).json({ error: 'Status is required' });
      return;
    }

    const updatedOrder = await updateOrderStatus(orderId, statusData);
    res.status(200).json({ message: 'Order status updated successfully', order: updatedOrder });
  } catch (error) {
    res.status(500).send('Error updating order status');
  }
};
