# Architecture Diagrams

This folder contains all system architecture diagrams for the Personal Finance App.

## 📊 Available Diagrams

1. **[System Overview](01-system-overview.md)** - Complete system architecture
2. **[Data Flow](02-data-flow.md)** - Request/response sequence
3. **[Database Schema](03-database-schema.md)** - Entity relationships
4. **[Temporal Workflow](04-temporal-workflow.md)** - Workflow architecture
5. **[Component Interaction](05-component-interaction.md)** - Component dependencies
6. **[Deployment](06-deployment.md)** - Production infrastructure
7. **[Security Layers](07-security-layers.md)** - Security architecture
8. **[Monitoring](08-monitoring.md)** - Observability stack

## 🎨 Viewing in Obsidian

### Enable Better Rendering

1. **Go to Settings** → Appearance → CSS snippets
2. **Enable** `mermaid-full-width.css`
3. Diagrams will now render full-width without horizontal scrolling

### Alternative: Use Mermaid Plugin

If diagrams still don't look good:
1. Install the **"Mermaid Tools"** plugin from Community Plugins
2. It provides better Mermaid rendering than the built-in support

## 🔄 Converting to Excalidraw

To convert any diagram to Excalidraw for editing:

1. **Export as SVG**:
   - Click the diagram
   - Right-click → "Export as SVG"
   
2. **Import to Excalidraw**:
   - Create new Excalidraw drawing
   - Drag the SVG file into it
   - Right-click → "Convert to Excalidraw"

## 📝 Editing Diagrams

Each `.md` file contains the Mermaid source code. To edit:
1. Open the file in edit mode
2. Modify the Mermaid code
3. Preview will update automatically

## 🎯 Tech Stack Shown

- **Frontend**: Flutter + Ferry + Drift
- **Backend**: Go + chi + gqlgen
- **Database**: PostgreSQL 16 with UUIDv7
- **Workflows**: Temporal
- **Auth**: Firebase Auth
- **Monitoring**: OpenTelemetry + Prometheus