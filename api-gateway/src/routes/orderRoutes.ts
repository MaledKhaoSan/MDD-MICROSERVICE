import { Router } from 'express';
import { getOrderById, patchOrderStatus, postOrder } from '../controllers/orderController';

const router = Router();

router.get('/:id', getOrderById);
router.post('/', postOrder);
router.patch('/status/:id', patchOrderStatus)

export default router;
