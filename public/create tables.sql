CREATE TABLE IF NOT EXISTS classes(
	id SERIAL PRIMARY KEY,
	fk_parent INT,
	name  VARCHAR(50) UNIQUE,
	CONSTRAINT classes_classes FOREIGN KEY(fk_parent) REFERENCES classes(id)
);
CREATE TABLE IF NOT EXISTS roles(
	id SERIAL PRIMARY KEY,
	userRole VARCHAR(30) unique NOT NULL
);
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	userName  VARCHAR(50) NOT NULL,
	surname  VARCHAR(50) NOT NULL,
	patronymic  VARCHAR(50) NOT NULL,
	login  VARCHAR(50) unique NOT NULL,
	userPassword  VARCHAR(20) NOT NULL,
	fk_role  INT NOT NULL,
	CONSTRAINT users_roles FOREIGN KEY(fk_role) REFERENCES roles(id)
);
CREATE TABLE IF NOT EXISTS equipments(
	id SERIAL PRIMARY KEY,
	fk_class INT NOT NULL,
	fk_user INT,
	inventoryNumber  VARCHAR(10) NOT NULL unique,
	equipmentName VARCHAR(50) NOT NULL UNIQUE,
	status INT NOT NULL,
	CONSTRAINT equipments_classes FOREIGN KEY(fk_class) REFERENCES classes(id),
	CONSTRAINT equipments_users FOREIGN KEY(fk_user) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS inventoryEvents(
	id SERIAL PRIMARY KEY,
	fk_equipment INT NOT NULL,
	fk_user INT NOT NULL,
	actionEvent VARCHAR(50) NOT NULL,
	date  date NOT NULL,
	CONSTRAINT inventoryEvents_users FOREIGN KEY(fk_user) REFERENCES users(id),
	CONSTRAINT inventoryEvents_equipments FOREIGN KEY(fk_equipment) REFERENCES equipments(id)
);
 
 
 
 
 
 

