/*
Let's discuss the pros and cons of the provided design, which follows the Abstract Factory design pattern for creating different types of report generators.

### Pros:

1. **Abstraction and Encapsulation:**
  - The Abstract Factory design pattern encapsulates the creation of related objects, providing a higher level of abstraction. Clients interact with the factory, not with the specific implementations of the products.

2. **Flexibility and Extensibility:**
  - The design allows for easy addition of new report generator types (e.g., PDF, CSV, Excel) without modifying client code. You can create new factories and products by implementing the necessary interfaces.

3. **Separation of Concerns:**
  - The client code is isolated from the specifics of object creation. It focuses on using the generated reports without being burdened with the creation logic.

4. **Consistent Product Creation:**
  - The factories ensure that the created report generators adhere to a specific interface, maintaining a consistent way to generate reports regardless of the type.

5. **Easy Configuration:**
  - Clients can request a specific type of report generator by providing a report type string. This simplifies the configuration and instantiation process.

6. **Potential for Dependency Injection:**
  - The design facilitates dependency injection, allowing clients to receive an AbstractFactory instance through constructor or method parameters.

### Cons:

1. **Complexity for Simple Scenarios:**
  - The Abstract Factory pattern might introduce unnecessary complexity for simple applications with few product types, making it over-engineered.

2. **Increased Code Volume:**
  - Implementing the abstract factory and concrete factory classes can lead to a larger codebase, which might be excessive for small-scale projects.

3. **Tight Coupling with Factories:**
  - The client code becomes tightly coupled to the specific factory implementations, making it challenging to switch between different factories.

4. **Limited Product Variability:**
  - Extending the product family (e.g., adding a new report generator type) requires modifications to the existing factories, potentially impacting the stability and maintenance.

5. **Runtime Efficiency:**
  - The Abstract Factory pattern may introduce a slight runtime overhead due to the abstraction layers involved in creating objects.

The suitability of the Abstract Factory pattern depends on the project's complexity, scalability requirements, and potential for future product family expansion. For simple applications, a simpler creational pattern like the Factory Method might be more appropriate. Always choose the design pattern that best fits the specific needs of your project.
*/
package creational

import "fmt"

type ReportGenerator interface {
	GenerateReport() error
}

type PDFGenerator struct{}

func (p *PDFGenerator) GenerateReport() error {
	fmt.Println("PDF report generated.")
	return nil
}

type CSVGenerator struct{}

func (c *CSVGenerator) GenerateReport() error {
	fmt.Println("CSV report generated.")
	return nil
}

type ExcelGenerator struct{}

func (e *ExcelGenerator) GenerateReport() error {
	fmt.Println("Excel report generated.")
	return nil
}

type AbstractFactory interface {
	CreateReportGenerator() ReportGenerator
}

type PDFGeneratorFactory struct{}

func (p *PDFGeneratorFactory) CreateReportGenerator() ReportGenerator {
	return &PDFGenerator{}
}

type CSVGeneratorFactory struct{}

func (c *CSVGeneratorFactory) CreateReportGenerator() ReportGenerator {
	return &CSVGenerator{}
}

type ExcelGeneratorFactory struct{}

func (e *ExcelGeneratorFactory) CreateReportGenerator() ReportGenerator {
	return &ExcelGenerator{}
}

func GetReportGeneratorFactory(reportType string) AbstractFactory {
	switch reportType {
	case "pdf":
		return &PDFGeneratorFactory{}
	case "csv":
		return &CSVGeneratorFactory{}
	case "excel":
		return &ExcelGeneratorFactory{}
	default:
		return nil
	}
}
