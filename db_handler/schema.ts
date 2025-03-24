import { date, decimal, integer, json, pgTable, text, timestamp, varchar } from "drizzle-orm/pg-core";

export const tenants = pgTable("tenants", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    name: text("name"),
    email: text("email"),
    address: text("address"),
    image: text("image"),
    properties: json("properties")
});

export const users = pgTable("users", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    name: text("name"),
    email: text("email").unique(),
    emailVerifiedAt: timestamp("email_verified_at", { mode: "date"}),
    image: text("image"),
    password: text("password"),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    properties: json("properties")
});

export const listings = pgTable("listings", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    name: text("name"),
    type: text("type"),
    purpose: text("purpose"),
    status: text("status"),
    address: text("address"),
    city: text("city"),
    region: text("region"),
    zip: text("zip"),
    community: text("community"),
    coordinates: json("coordinates"),
    price: decimal("price"),
    currency: text("currency"),
    price_type: text("price_type"),
    taxes: decimal("taxes"),
    fees: decimal("fees"),
    size: decimal("size"),
    bedrooms: integer("bedrooms"),
    bathrooms: integer("bathrooms"),
    floor: decimal("floor"),
    year_built: text("year_built"),
    parking: integer("parking"),
    furnishing_status: text("furnishing_status"),
    indoor_features: json("indoor_features"),
    outdoor_features: json("outdoor_features"),
    community_features: json("community_features"),
    utilities: json("utilities"),
    agent_name: text("agent_name"),
    agent_contact: text("agent_contact"),
    agency_name: text("agency_name"),
    description: text("description"),
    nearby_landmarks: json("nearby_landmarks"),
    availability: text("availability"),
    condition: text("condition"),
    tags: json("tags"),
    images: json("images"),
    videos: json("videos"),
    properties: json("properties")
});

export const customers = pgTable("customers", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    name: text("name"),
    company: text("company"),
    email: text("email"),
    phone: text("phone"),
    status: text("status"),
    industry: text("industry"),
    properties: json("properties")
});

export const deals = pgTable("deals", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    listingId: text("listing_id")
      .notNull()
      .references(() => listings.id, { onDelete: "cascade" }),
    customerId: text("customer_id")
      .notNull()
      .references(() => customers.id, { onDelete: "cascade" }),
    amount: decimal("amount"),
    stage: text("stage"),
    close_date: date("close_date"),
    properties: json("properties")
});

export const tasks = pgTable("tasks", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    title: text("title"),
    description: text("description"),
    due_date: date("due_date"),
    priority: text("priority"),
    status: text("status"),
    properties: json("properties")
});

export const report = pgTable("reports", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    title: text("title"),
    description: text("description"),
    data: json("data"),
    properties: json("properties")
});

export const notifications = pgTable("noifications", {
    id: text("id")
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
    tenantId: text("tenant_id")
      .notNull()
      .references(() => tenants.id, { onDelete: "cascade" }),
    data: json("data"),
    properties: json("properties"),
    type: text("type")
});