-- name: GetOrderInfo :many
SELECT p.name as product_name, p.id as product_id, o.id as order_id, oi.quantity as quantity, s.name as shelve_name, ps.is_primary as shelve_is_primary
FROM orders o
         INNER JOIN order_items oi on o.id = oi.order_id
         INNER JOIN products p on p.id = oi.product_id
         INNER JOIN products_shelves ps on p.id = ps.product_id
         INNER JOIN public.shelves s on s.id = ps.shelves_id
WHERE o.id = ANY($1::bigint[])
ORDER BY o.id, p.id, ps.is_primary DESC;