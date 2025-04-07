import { Request, Response } from 'express';
import { AddItem, ReduceItem } from '../services/apiService';

import { Inventory } from '../models/inventoryModel';

export const AddQuantityItem = async (req: Request, res: Response): Promise<void> => {
  try {
    const inventoryId = req.params.id;
    const quantity: Inventory = req.body;

    if (!inventoryId) {
      res.status(400).json({ error: 'Inventory ID is required' });
      return;
    }

    if (!quantity.quantity_change || typeof quantity.quantity_change !== 'number' || quantity.quantity_change <= 0) {
      res.status(400).json({ error: 'Quantity change must be a positive number' });
      return;
    }

    const addItem = await AddItem(inventoryId, quantity);
    res.status(200).json({ message: 'Item added to inventory successfully', item: addItem });
  }
  catch (error) {
    res.status(500).send('Error adding item to inventory');
  }
};

export const ReduceQuantityItem = async (req: Request, res: Response): Promise<void> => {
  try {
    const inventoryId = req.params.id;
    const quantity: Inventory = req.body;

    if (!inventoryId) {
      res.status(400).json({ error: 'Inventory ID is required' });
      return;
    }

    if (!quantity.quantity_change || typeof quantity.quantity_change !== 'number' || quantity.quantity_change <= 0) {
      res.status(400).json({ error: 'Quantity change must be a positive number' });
      return;
    }

    const addItem = await ReduceItem(inventoryId, quantity);
    res.status(200).json({ message: 'Item reduced to inventory successfully', item: addItem });
  }
  catch (error) {
    res.status(500).send('Error adding item to inventory');
  }
};

