#ifndef SEARCHWINDOW_H
#define SEARCHWINDOW_H

#include <QMainWindow>

namespace Ui {
class SearchWindow;
}

class SearchWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit SearchWindow(QWidget *parent = 0);
    ~SearchWindow();
signals:
    void plsSearchStudName(QString data);
    void plsSearchStudSur(QString data);
    void plsSearchCours(QString data);
private slots:
    void on_snsearchB_clicked();

    void on_sssearchB_clicked();

    void on_cnsearchB_clicked();

private:
    Ui::SearchWindow *ui;
};

#endif // SEARCHWINDOW_H
