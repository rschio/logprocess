CREATE TABLE IF NOT EXISTS requests (
	id varchar(255) NOT NULL,
	method varchar(255) NOT NULL,
	uri varchar(255) NOT NULL,
	url varchar(255) NOT NULL,
	size BIGINT NOT NULL,
	querystring varchar(255) NOT NULL,
	header_accept varchar(255) NOT NULL,
	header_host varchar(255) NOT NULL,
	header_user_agent varchar(255) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS responses (
	id varchar(255) NOT NULL,
	status BIGINT NOT NULL,
	size BIGINT NOT NULL,
	content_length BIGINT NOT NULL,
	via varchar(255) NOT NULL,
	connection varchar(255) NOT NULL,
	access_control_allow_credentials varchar(255) NOT NULL,
	access_control_allow_origin varchar(255) NOT NULL,
	content_type varchar(255) NOT NULL,
	server varchar(255) NOT NULL,
	PRIMARY KEY (id)	
);

CREATE TABLE IF NOT EXISTS routes (
	id varchar(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	hosts varchar(255) NOT NULL,
	methods varchar(255) NOT NULL,
	paths varchar(255) NOT NULL,
	preserve_host BOOLEAN NOT NULL,
	protocols varchar(255) NOT NULL,
	regex_priority BIGINT NOT NULL,
	service_id varchar(255) NOT NULL,
	strip_path BOOLEAN NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS services (
	id varchar(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	host varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	path varchar(255) NOT NULL,
	port BIGINT NOT NULL,
	protocol varchar(255) NOT NULL,
	read_timeout BIGINT NOT NULL,
	write_timeout BIGINT NOT NULL,
	connect_timeout BIGINT NOT NULL,
	retries BIGINT NOT NULL,
	PRIMARY KEY (id)	
);


CREATE TABLE IF NOT EXISTS records (
	id BIGINT AUTO_INCREMENT NOT NULL,
	consumer_id varchar(255) NOT NULL,
	upstream_uri varchar(255) NOT NULL,
	response_id varchar(255) NOT NULL,
	request_id varchar(255) NOT NULL,
	route_id varchar(255) NOT NULL,	
	service_id varchar(255) NOT NULL,
	proxy_latency BIGINT NOT NULL,
	gateway_latency BIGINT NOT NULL,
	request_latency BIGINT NOT NULL,
	client_ip varchar(50) NOT NULL,
	started_at TIMESTAMP NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (response_id) REFERENCES responses(id),
	FOREIGN KEY (request_id) REFERENCES requests(id),
	FOREIGN KEY (route_id) REFERENCES routes(id),
	FOREIGN KEY (service_id) REFERENCES services(id)
);
