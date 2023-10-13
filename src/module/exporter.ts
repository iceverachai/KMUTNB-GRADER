
type ImportData = {
    header: Array<string>
    row: Array<any>
}

export const ExportCSV = (data : ImportData) => {
    const csvContent = [
        data.header.join(','), // Create CSV header
        ...data.row.map((item:any) => Object.values(item).join(',')) // Create CSV rows
      ].join('\n');
    const BOM = '\uFEFF';
    
    const blob = new Blob([BOM+csvContent], { type: 'text/csv; charset=utf-8;' });

    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = 'output.csv';
    link.click();
}
