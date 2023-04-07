WITH order_product_data AS (
	SELECT
		oi.order_id AS order_id,
		JSONB_AGG(
            JSON_BUILD_OBJECT(
                'id', oi.id,
                'order_id', oi.order_id,
                'product_id', oi.product_id,
            )
        ) AS order_products
		
	FROM order_products AS oi
    JOIN product AS p ON oi.product_id = p.id
	WHERE oi.order_id = '05102f47-8dbe-4c80-b8db-0a00d0ad2c28'
	GROUP BY oi.order_id
)
SELECT
	o.id, 
	o.client_id, 
	c.id, 
	c.first_name,
	c.last_name,
	c.phone_number,
	CAST(c.created_at::timestamp AS VARCHAR),
	CAST(c.updated_at::timestamp AS VARCHAR),
	COALESCE(o.price, 0),
	COALESCE(o.status, ''),
	CAST(o.created_at::timestamp AS VARCHAR),
	CAST(o.updated_at::timestamp AS VARCHAR),
    op.order_products
FROM "orders" AS o
JOIN client AS c ON c.id = o.client_id
JOIN order_product_data AS op ON op.order_id = o.id 
WHERE o.id = '05102f47-8dbe-4c80-b8db-0a00d0ad2c28';