package query

// SQL for RAG vector search schema
const (
	// UseTestDatabaseSQL switches to the target database
	UseTestDatabaseSQL = "USE test;"

	// CreateEmbeddedDocumentsTableSQL creates the table for storing documents and their embeddings
	CreateEmbeddedDocumentsTableSQL = `
CREATE TABLE embedded_documents (
    id        INT       PRIMARY KEY,
    -- Column to store the original content of the document.
    document  TEXT,
    -- Column to store the vector representation of the document.
    embedding VECTOR(3072)
);`

	// InsertSampleDocumentsSQL inserts sample rows into embedded_documents
	InsertSampleDocumentsSQL = `
INSERT INTO embedded_documents
VALUES
    (1, 'dog', '[1,2,1]'),
    (2, 'fish', '[1,2,4]'),
    (3, 'tree', '[1,0,0]');`

	// SelectTop3ByCosineSQL selects the top 3 nearest documents by cosine distance
	SelectTop3ByCosineSQL = `
SELECT id, document, vec_cosine_distance(embedding, '[1,2,3]') AS distance
FROM embedded_documents
ORDER BY distance
LIMIT 3;`
)

// ToDo Insert operation

// ToDo search Operation

// ToDo add index Operation
