import { Router } from 'express';
import { AddQuantityItem, ReduceQuantityItem } from '../controllers/inventoryController';
const router = Router();

router.patch('/add/:id', AddQuantityItem)
router.patch('/reduce/:id', ReduceQuantityItem)

export default router;
